package cmd

import (
	"fmt"
	"log"

	"github.com/justinfinch/gauge/conf"
	"github.com/spf13/cobra"
)

// NewServeCommand creates and sets up the serve command
func NewServeCommand() *cobra.Command {
	serveCmd := cobra.Command{
		Use: "serve",
		Run: serveRun,
	}

	// this is where we will configure everything!
	serveCmd.Flags().IntP("port", "p", 80, "the port to do things on")

	return &serveCmd
}

func serveRun(cmd *cobra.Command, args []string) {
	fmt.Println("--- serve command---")

	config, err := conf.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	logger, err := conf.ConfigureLogging(config)
	if err != nil {
		log.Fatal("Failed to configure logging: " + err.Error())
	}

	logger.Infof("Starting with config: %+v", config)

	api, err := api.NewAPI(logger, config)
	if err != nil {
		log.Fatal("Failed to start api server: " + err.Error())
	}
	api.Start()

}
