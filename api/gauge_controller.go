package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/justinfinch/gauge/data"
	"github.com/justinfinch/gauge/model"
	"github.com/labstack/echo"
)

type gaugeCreate struct {
	Name string
}

func (api *API) registerGaugeRoutes() {
	api.echo.GET("tenant/:tenantId/gauges", searchGauges(api.log, api.db))
	api.echo.GET("tenant/:tenantId/gauges/:id", getGauge(api.log))
	api.echo.POST("tenant/:tenantId/gauges", createGauge(api.log, api.db))
}

func searchGauges(log *logrus.Entry, db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("Searching gauges")

		gaugeRepo, err := data.NewGaugeRepo(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		tenatnID := c.Param("tenantId")
		gauges, err := gaugeRepo.GetAll(tenatnID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		return c.JSON(http.StatusOK, gauges)
	}
}

func getUserGauges(log *logrus.Entry) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("Getting user gauges")

		return c.JSON(http.StatusOK, map[string]string{
			"description": "a boiler plate project",
			"name":        "gauge",
		})
	}
}

func getGauge(log *logrus.Entry) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("Getting gauge by id")

		id := c.Param("id")
		return c.JSON(http.StatusOK, map[string]string{
			"description": "a boiler plate project",
			"name":        "gauge",
			"id":          id,
		})
	}
}

func createGauge(log *logrus.Entry, db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := new(gaugeCreate)
		if err := c.Bind(request); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		//TODO: Add request validation

		log.WithFields(logrus.Fields{
			"name": request.Name,
		}).Debug("Creating gauge")

		tenatnID := c.Param("tenantId")
		gauge, err := model.NewGauge(request.Name, tenatnID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		tx := db.Begin()

		gaugeRepo, err := data.NewGaugeRepo(tx)
		if err != nil {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, nil)
		}

		err = gaugeRepo.Save(gauge)
		if err != nil {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, nil)
		}

		tx.Commit()
		return c.JSON(http.StatusCreated, gauge)
	}
}
