package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/justinfinch/gauge/conf"
	"github.com/justinfinch/gauge/gauges"
)

//API Serves routes
type API struct {
	log    *logrus.Entry
	config *conf.Config
	db     *gorm.DB
}

//NewAPI creates a new API
func NewAPI(log *logrus.Entry, config *conf.Config) (*API, error) {
	api := &API{
		log:    log.WithField("component", "api"),
		config: config,
	}

	const addr = "postgresql://gaugeapp@localhost:26257/gaugedb?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	api.db = db

	gaugeRepo := gauges.NewGormGaugeRepo(api.db)
	gaugeService := gauges.NewService(gaugeRepo, api.log)

	return api, nil
}

//Start starts the api server
func (api *API) Start() {
	defer api.db.Close()
	api.echo.Start(fmt.Sprintf(":%d", api.config.Port))
	api.log.Info("Start method exited")
}
