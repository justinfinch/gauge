package api

import (
	"net/http"
	"strconv"

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
	api.echo.GET("org/:orgId/gauges", searchGauges(api.log, api.db))
	api.echo.GET("org/:orgId/gauges/:id", getGauge(api.log, api.db))
	api.echo.POST("org/:orgId/gauges", createGauge(api.log, api.db))
}

func searchGauges(log *logrus.Entry, db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("Searching gauges")

		gaugeRepo, err := data.NewGaugeRepo(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		tenatnID := c.Param("orgId")
		gauges, err := gaugeRepo.GetAll(tenatnID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		return c.JSON(http.StatusOK, gauges)
	}
}

func getGauge(log *logrus.Entry, db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("Getting gauge by id")

		gaugeRepo, err := data.NewGaugeRepo(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		orgID := c.Param("orgId")
		gauge, err := gaugeRepo.Get(orgID, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		return c.JSON(http.StatusOK, gauge)
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

		tenatnID := c.Param("orgId")
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
