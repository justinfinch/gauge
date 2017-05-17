package model

import (
	"github.com/google/uuid"
)

//Gauge represents something that can have a value read from it creating a reading
type Gauge struct {
	ID   int64 `gorm:"primary_key"`
	UUID string
	Name string
}

//NewGauge creates a new gauge struct given a name
func NewGauge(name string) *Gauge {
	gaugeUUID, _ := uuid.NewRandom()
	gauge := &Gauge{
		UUID: gaugeUUID.String(),
		Name: name,
	}

	return gauge
}

func (gauge *Gauge) Save() error {

	return nil
}
