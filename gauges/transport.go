package gauges

import (
	"log"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	createGaugeHandler := kithttp.NewServer()
		makeCreateGaugeEndpoint(s),
		decodeCreateGaugeRequest,
		encodeResponse
	)


}

// func (api *API) registerGaugeRoutes() {
// 	api.echo.GET("org/:orgId/gauges", searchGauges(api.log, api.db))
// 	api.echo.GET("org/:orgId/gauges/:id", getGauge(api.log, api.db))
// 	api.echo.POST("org/:orgId/gauges", createGauge(api.log, api.db))
// }
