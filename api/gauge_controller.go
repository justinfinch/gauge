package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/justinfinch/gauge/model"
	"github.com/labstack/echo"
)

type gaugeCreate struct {
	Name string
}

func (api *API) registerGaugeRoutes() {
	api.echo.GET("user/:userId/gauges", getUserGauges(api.log))
	api.echo.GET("/gauges", searchGauges(api.log))
	api.echo.GET("/gauges/:id", getGauge(api.log))
	api.echo.POST("/gauges", createGauge(api.log))
}

func searchGauges(log *logrus.Entry) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("Searching gauges")
		return c.JSON(http.StatusOK, map[string]string{
			"description": "a boiler plate project",
			"name":        "gauge",
		})
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

func createGauge(log *logrus.Entry) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := new(gaugeCreate)
		if err := c.Bind(request); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		//TODO: Add request validation

		log.WithFields(logrus.Fields{
			"name": request.Name,
		}).Debug("Creating gauge")

		gauge := model.NewGauge(request.Name)

		return c.JSON(http.StatusCreated, gauge)
	}
}
