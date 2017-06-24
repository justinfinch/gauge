package gauges

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//CREATE GAUGE ENDPOINT

type createGaugeRequest struct {
	Name  string
	OrgID string
}

type createGaugeResponse struct {
	Gauge *Gauge `json:"gauge,omitempty"`
	Err   error  `json:"error,omitempty"`
}

func makeCreateGaugeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createGaugeRequest)

		gauge, err := NewGauge(req.Name, req.OrgID)
		if err != nil {
			return createGaugeResponse{Err: err}, err
		}

		gauge, err = s.CreateNewGauge(gauge)
		if err != nil {
			return createGaugeResponse{Err: err}, err
		}

		return createGaugeResponse{Gauge: gauge}, nil
	}
}

//FIND GAUGE ENDPOINT

type findGaugeRequest struct {
	ID    uint
	OrgID string
}

type findGaugeResponse struct {
	Gauge *Gauge `json:"gauge,omitempty"`
	Err   error  `json:"error,omitempty"`
}

func makeFindGaugeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(findGaugeRequest)

		gauge, err := s.FindGauge(req.ID, req.OrgID)
		if err != nil {
			return findGaugeResponse{Err: err}, err
		}

		return findGaugeResponse{Gauge: gauge}, nil
	}
}

//SEARCH GAUGE ENDPOINT

type searchGaugeRequest struct {
	OrgID string
}

type searchGaugeResponse struct {
	Gauges *[]Gauge `json:"gauges,omitempty"`
	Err    error    `json:"error,omitempty"`
}

func searchFindGaugeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(searchGaugeRequest)

		gauges, err := s.SearchGauges(req.OrgID)
		if err != nil {
			return searchGaugeResponse{Err: err}, err
		}

		return searchGaugeResponse{Gauges: gauges}, nil
	}
}
