package data

import "github.com/jinzhu/gorm"
import "github.com/justinfinch/gauge/model"

//GaugeRepo is used to save gauge data to the database
type GaugeRepo struct {
	DB *gorm.DB
}

//NewGaugeRepo creates a new gauge repo
func NewGaugeRepo(db *gorm.DB) (*GaugeRepo, error) {
	repo := &GaugeRepo{
		DB: db,
	}
	return repo, nil
}

//Save creates or updates a gauge record
func (repo *GaugeRepo) Save(gauge *model.Gauge) error {
	if repo.DB.NewRecord(gauge) {
		repo.DB.Create(gauge)
	}

	return nil
}

//GetAll gets all guages
func (repo *GaugeRepo) GetAll(orgID string) (*[]model.Gauge, error) {
	var gauges []model.Gauge
	repo.DB.Where(&model.Gauge{OrgID: orgID}).Find(&gauges)

	return &gauges, nil
}

//Get gets a guage by ID
func (repo *GaugeRepo) Get(orgID string, id int64) (*model.Gauge, error) {
	var gauge model.Gauge
	repo.DB.Where(&model.Gauge{OrgID: orgID}).First(&gauge, id)
	return &gauge, nil
}
