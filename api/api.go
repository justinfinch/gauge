package api

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/justinfinch/gauge/conf"
	"github.com/labstack/echo"
)

//API Serves routes
type API struct {
	log    *logrus.Entry
	config *conf.Config
	echo   *echo.Echo
	db     *gorm.DB
}

//NewAPI creates up a new API
func NewAPI(log *logrus.Entry, config *conf.Config) (*API, error) {
	api := &API{
		log:    log.WithField("component", "api"),
		config: config,
	}

	e := echo.New()
	e.Use(api.loggingMiddleware)
	api.echo = e

	api.registerGaugeRoutes()

	return api, nil
}

//Start starts the api server
func (api *API) Start() {
	const addr = "postgresql://gaugeapp@localhost:26257/gauge?sslmode=disable"
	//db, err := gorm.Open("postgres", addr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()

	api.echo.Start(fmt.Sprintf(":%d", api.config.Port))
	api.log.Info("Start method exited")
}

func (api *API) loggingMiddleware(f echo.HandlerFunc) echo.HandlerFunc {
	// return a HandlerFunc
	return func(ctx echo.Context) error {
		req := ctx.Request()
		// add some default fields to the logger ~ on all messages
		reqID, _ := uuid.NewRandom()
		logger := api.log.WithFields(logrus.Fields{
			"method":     req.Method,
			"path":       req.URL.Path,
			"request_id": reqID.String(),
		})
		ctx.Set("logger", logger)
		startTime := time.Now()

		defer func() {
			rsp := ctx.Response()
			// at the end we will want to log a few more interesting fields
			logger.WithFields(logrus.Fields{
				"status_code":  rsp.Status,
				"runtime_nano": time.Since(startTime).Nanoseconds(),
			}).Debug("Finished request")
		}()

		// now we will log out that we have actually started the request
		logger.WithFields(logrus.Fields{
			"user_agent":     req.UserAgent(),
			"content_length": req.ContentLength,
		}).Debug("Starting request")

		err := f(ctx)
		if err != nil {
			ctx.Error(err)
		}

		return err
	}
}
