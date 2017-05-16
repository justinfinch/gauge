package model

import (
	"github.com/google/uuid"
)

//Gauge represents something that can have a value read from it creating a reading
type Gauge struct {
	ID   uuid.UUID `gorm:"primary_key"`
	Name string
}

//NewGauge creates a new gauge struct given a name
func NewGauge(name string) *Gauge {
	gaugeID, _ := uuid.NewRandom()
	gauge := &Gauge{
		ID:   gaugeID,
		Name: name,
	}

	return gauge
}

func (gauge *Gauge) Save() error {

	return nil
}
