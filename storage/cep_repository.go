package storage

import (
	"modules-app/controllers"

	"gorm.io/gorm"
)

type CepRepository struct {
	db *gorm.DB
}

func NewCepRepository(db *gorm.DB) *CepRepository {
	return &CepRepository{db: db}
}

func (r *CepRepository) GetAllCEPS() ([]controllers.CEP, error) {
	var ceps []controllers.CEP
	if err := r.db.Find(&ceps).Error; err != nil {
		return nil, err
	}
	return ceps, nil
}

func (r *CepRepository) GetCEPByID(id string) (controllers.CEP, error) {
	var cep controllers.CEP
	if err := r.db.First(&cep, "id = ?", id).Error; err != nil {
		return controllers.CEP{}, err
	}
	return cep, nil
}

func (r *CepRepository) CreateNewCEP(cep controllers.CEP) error {
	result := r.db.Create(&cep)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
