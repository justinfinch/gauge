package cmd

import (
	"log"

	"github.com/jinzhu/gorm"

	//Import of postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/justinfinch/gauge/model"
	"github.com/spf13/cobra"
)

// NewDbCommand creates and sets up the db command
func NewMigrateCommand() *cobra.Command {
	migrateCmd := cobra.Command{
		Use: "db-migrate",
		Run: migrateRun,
	}

	return &migrateCmd

}

func migrateRun(cmd *cobra.Command, args []string) {
	const addr = "postgresql://gaugeapp@localhost:26257/gaugedb?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Automatically create the "Gauges" table based on the Gauge model.
	db.AutoMigrate(&model.Gauge{})

}
