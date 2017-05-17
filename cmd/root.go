package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewRootCommand creates and sets up the root command
func NewRootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "gauge",
		Run: rootRun,
	}

	// this is where we will configure everything!
	rootCmd.PersistentFlags().StringP("loglevel", "l", "info", "the log level")

	serveCmd := NewServeCommand()
	rootCmd.AddCommand(serveCmd)

	migrateCmd := NewMigrateCommand()
	rootCmd.AddCommand(migrateCmd)

	return &rootCmd
}

func rootRun(cmd *cobra.Command, args []string) {
	fmt.Println("--- root command---")

}
