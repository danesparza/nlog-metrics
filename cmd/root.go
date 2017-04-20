package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile               string
	problemWithConfigFile bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "nlog-metrics",
	Short: "A brief description of your application",
	Long: `nlog-metrics is a service to gather metrics from an NLog database 
and send that information to InfluxDB`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	//	Set our defaults
	viper.SetDefault("sqlserver.server", "servernamehere")
	viper.SetDefault("sqlserver.database", "system_logging")
	viper.SetDefault("sqlserver.user", "user")
	viper.SetDefault("sqlserver.password", "password")
	viper.SetDefault("influxdb.server", "localhost:8086")
	viper.SetDefault("influxdb.database", "applogs")
	viper.SetDefault("influxdb.measurement", "metrics")

	//	Set our config search locations
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AddConfigPath(".")      // also look in the working directory
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in
	// otherwise, make note that there was a problem
	if err := viper.ReadInConfig(); err != nil {
		problemWithConfigFile = true
	}
}
