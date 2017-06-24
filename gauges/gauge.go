package gauges

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//Gauge represents something that can have a value read from it creating a reading
type Gauge struct {
	gorm.Model
	UUID     string `gorm:"not null"`
	Name     string `gorm:"not null"`
	OrgID 	 string `gorm:"not null"`
}

//NewGauge creates a new gauge struct given a name
func NewGauge(name string, orgID string) (*Gauge, error) {
	gaugeUUID, _ := uuid.NewRandom()
	gauge := &Gauge{
		UUID:     gaugeUUID.String(),
		Name:     name,
		OrgID:    orgID,
	}

	return gauge, nil
}
