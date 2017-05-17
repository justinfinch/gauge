package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//Gauge represents something that can have a value read from it creating a reading
type Gauge struct {
	gorm.Model
	UUID string
	Name string
}

//NewGauge creates a new gauge struct given a name
func NewGauge(name string) (*Gauge, error) {
	gaugeUUID, _ := uuid.NewRandom()
	gauge := &Gauge{
		UUID: gaugeUUID.String(),
		Name: name,
	}

	return gauge, nil
}
