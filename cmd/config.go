package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	jsonConfig bool
	yamlConfig bool
)

var yamlDefault = []byte(`
sqlserver:
  server: servernamehere
  database: system_logging
  user: username
  password: password
influxdb:
  server: http://localhost:8086
  database: applogs
  measurement: metrics
`)

var jsonDefault = []byte(`{
	"sqlserver": {
		"server": "servernamehere",
		"database": "system_logging",
		"user": "username",
		"password": "password"
	},
	"influxdb": {
		"server": "http://localhost:8086",
		"database": "applogs",
		"measurement": "metrics"
	}
}`)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Prints default server configuration files",
	Long: `Use this to create a default configuration file for the nlog-metrics server  

Example: 

nlog-metrics config > config.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		if jsonConfig {
			fmt.Printf("%s", jsonDefault)
		} else if yamlConfig {
			fmt.Printf("%s", yamlDefault)
		}
	},
}

func init() {
	RootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolVarP(&jsonConfig, "json", "j", false, "Create a JSON configuration file")
	configCmd.Flags().BoolVarP(&yamlConfig, "yaml", "y", true, "Create a YAML configuration file")

}
