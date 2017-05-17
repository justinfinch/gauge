package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//Gauge represents something that can have a value read from it creating a reading
type Gauge struct {
	gorm.Model
	UUID     string `gorm:"not null"`
	Name     string `gorm:"not null"`
	TenantID string `gorm:"not null"`
}

//NewGauge creates a new gauge struct given a name
func NewGauge(name string, tenantID string) (*Gauge, error) {
	gaugeUUID, _ := uuid.NewRandom()
	gauge := &Gauge{
		UUID:     gaugeUUID.String(),
		Name:     name,
		TenantID: tenantID,
	}

	return gauge, nil
}
