package gauges

import "github.com/Sirupsen/logrus"

//Service is the interface provides methods for managing gauges
type Service interface {
	CreateNewGauge(gauge *Gauge) (*Gauge, error)
	FindGauge(id uint, orgID string) (*Gauge, error)
	SearchGauges(orgID string) (*[]Gauge, error)
}

type service struct {
	gaugeRepo GaugeRepo
	log       *logrus.Entry
}

//NewService creates a new Gauge Service with nessasary dependencies
func NewService(gaugeRepo GaugeRepo, log *logrus.Entry) Service {
	svcLog := log.WithFields(logrus.Fields{
		"componet": "Gauge Service",
	})

	return &service{
		gaugeRepo: gaugeRepo,
		log:       svcLog,
	}
}

func (s *service) CreateNewGauge(gauge *Gauge) (*Gauge, error) {
	s.log.Debug("Creating gauge")

	err := s.gaugeRepo.Save(gauge)
	if err != nil {
		return gauge, err
	}

	return gauge, nil
}

func (s *service) FindGauge(id uint, orgID string) (*Gauge, error) {
	s.log.Debug("Finding gauge")

	gauge, err := s.gaugeRepo.Get(orgID, id)
	if err != nil {
		return nil, err
	}

	return gauge, nil
}

func (s *service) SearchGauges(orgID string) (*[]Gauge, error) {
	s.log.Debug("Searching gauges")

	gauges, err := s.gaugeRepo.GetAll(orgID)
	if err != nil {
		return nil, err
	}

	return gauges, nil
}
