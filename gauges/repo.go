package gauges

import "github.com/jinzhu/gorm"

type GaugeRepo interface {
	Save(gauge *Gauge) error
	GetAll(orgID string) (*[]Gauge, error)
	Get(orgID string, id uint) (*Gauge, error)
}

//GormGaugeRepo is used to save gauge data to the database
type GormGaugeRepo struct {
	DB *gorm.DB
}

//NewGormGaugeRepo creates a new gauge repo
func NewGormGaugeRepo(db *gorm.DB) (*GormGaugeRepo, error) {
	repo := &GormGaugeRepo{
		DB: db,
	}
	return repo, nil
}

//Save creates or updates a gauge record
func (repo *GormGaugeRepo) Save(gauge *Gauge) error {
	if repo.DB.NewRecord(gauge) {
		repo.DB.Create(gauge)
	}

	return nil
}

//GetAll gets all guages
func (repo *GormGaugeRepo) GetAll(orgID string) (*[]Gauge, error) {
	var gauges []Gauge
	repo.DB.Where(&Gauge{OrgID: orgID}).Find(&gauges)

	return &gauges, nil
}

//Get gets a guage by ID
func (repo *GormGaugeRepo) Get(orgID string, id uint) (*Gauge, error) {
	var gauge Gauge
	repo.DB.Where(&Gauge{OrgID: orgID}).First(&gauge, id)
	return &gauge, nil
}
