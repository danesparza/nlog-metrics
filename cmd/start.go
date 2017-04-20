package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"

	"github.com/danesparza/nlog-metrics/data"
	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the monitoring process",
	Long: `Starts the metrics gathering and sending process

Example:

nlog-metrics start`,
	Run: execute,
}

func init() {
	RootCmd.AddCommand(startCmd)
}

func execute(cmd *cobra.Command, args []string) {
	log.Println("[INFO] Initializing...")

	//	Trap program exit appropriately
	ctx, cancel := context.WithCancel(context.Background())
	sigch := make(chan os.Signal, 2)
	signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
	go handleSignals(sigch, ctx, cancel)

	//	If we have a config file, report it:
	if viper.ConfigFileUsed() != "" {
		log.Println("[INFO] Using config file:", viper.ConfigFileUsed())
	}

	//	Report our config options:
	log.Printf("[INFO] SQL server: %s\n", viper.GetString("sqlserver.server"))
	log.Printf("[INFO] --- database: %s\n", viper.GetString("sqlserver.database"))
	log.Printf("[INFO] --- user: %s\n", viper.GetString("sqlserver.user"))
	log.Printf("[INFO] --- password: **not shown**\n")

	log.Printf("[INFO] InfluxDB server: %s\n", viper.GetString("influxdb.server"))
	log.Printf("[INFO] --- database: %s\n", viper.GetString("influxdb.database"))
	log.Printf("[INFO] --- measurement: %s\n", viper.GetString("influxdb.measurement"))

	//	Initialize our database object:
	nlogDatabase := data.MSSqlDB{
		Server:   viper.GetString("sqlserver.server"),
		Database: viper.GetString("sqlserver.database"),
		User:     viper.GetString("sqlserver.user"),
		Password: viper.GetString("sqlserver.password")}

	//	Initialize connection to InfluxDB
	c, err := influx.NewHTTPClient(influx.HTTPConfig{Addr: viper.GetString("influxdb.server")})
	if err != nil {
		log.Fatal(err)
	}

	//	Loop
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(30 * time.Second):
			//	Check the NLog database
			nlogItems, err := nlogDatabase.GetMetrics()
			if err != nil {
				log.Fatalf("[ERROR] Error getting metrics: %v", err)
			}

			// Create a new point batch
			bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
				Database:  viper.GetString("influxdb.database"),
				Precision: "s",
			})
			if err != nil {
				log.Fatal(err)
			}

			//	For each item, log to InfluxDB
			for _, nlogItem := range nlogItems {

				//	Create the tags and field information for the item:
				tags := map[string]string{
					"log_level":   nlogItem.LogLevel,
					"application": nlogItem.Application}
				fields := map[string]interface{}{
					"count": nlogItem.Count}

				//	Create a new point with our tags and field data
				pt, err := influx.NewPoint(viper.GetString("influxdb.measurement"), tags, fields, time.Now())
				if err != nil {
					log.Fatal(err)
				}

				//	Add the point to the batch
				bp.AddPoint(pt)
			}

			// Write the batch
			if err := c.Write(bp); err != nil {
				log.Fatal(err)
			}

		}
	}

}

func handleSignals(sigch <-chan os.Signal, ctx context.Context, cancel context.CancelFunc) {
	select {
	case <-ctx.Done():
	case sig := <-sigch:
		switch sig {
		case os.Interrupt:
			log.Println("[INFO] SIGINT - Shutting down")
		case syscall.SIGTERM:
			log.Println("[INFO] SIGTERM - Shutting down")
		}
		cancel()
	}
}
