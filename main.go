package main

import (
	"log"
	"os"

	"github.com/danesparza/nlog-metrics/cmd"
	"github.com/hashicorp/logutils"
)

func main() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("INFO"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	cmd.Execute()
}
