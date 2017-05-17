package cmd

import "github.com/spf13/cobra"

// NewDbCommand creates and sets up the db command
func NewMigrateCommand() *cobra.Command {
	migrateCmd := cobra.Command{
		Use: "db",
		Run: migrateRun,
	}

	return &migrateCmd

}

func migrateRun(cmd *cobra.Command, args []string) {

}
