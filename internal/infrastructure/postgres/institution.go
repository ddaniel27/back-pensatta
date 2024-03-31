package postgres

import (
	"context"
	"pensatta/internal/core/domain"
	"pensatta/internal/infrastructure/postgres/models"

	"gorm.io/gorm"
)

type InstitutionRepository struct {
	db *gorm.DB
}

func NewInstitutionRepository(db *gorm.DB) *InstitutionRepository {
	return &InstitutionRepository{
		db: db,
	}
}

func (ir *InstitutionRepository) Create(ctx context.Context, institution domain.Institution) error {
	institutionModel := toInstitutionModel(institution)

	if err := ir.db.WithContext(ctx).Create(&institutionModel).Error; err != nil {
		return err
	}

	languageModel := models.LanguageModel{
		InstitutionID: institutionModel.ID,
		Value:         institution.Language,
	}

	if err := ir.db.WithContext(ctx).Create(&languageModel).Error; err != nil {
		return err
	}

	return nil
}

func (ir *InstitutionRepository) Get(ctx context.Context) ([]domain.Institution, error) {
	var institutions []domain.Institution

	if err := ir.db.WithContext(ctx).Raw(`
	SELECT i.*, COALESCE(l.value, 'es') AS language
	FROM pensatta_institution i
	LEFT JOIN pensatta_languages l ON i.id = l.institution_id`).
		Scan(&institutions).Error; err != nil {
		return nil, err
	}

	return institutions, nil
}

func (ir *InstitutionRepository) Delete(ctx context.Context, id uint64) error {
	if err := ir.db.WithContext(ctx).Delete(&models.InstitutionModel{}, id).Error; err != nil {
		return err
	}

	return nil
}

func toInstitutionModel(institution domain.Institution) models.InstitutionModel {
	return models.InstitutionModel{
		Name:     institution.Name,
		Email:    institution.Email,
		Country:  institution.Country,
		Province: institution.Province,
		City:     institution.City,
		Code:     institution.Code,
	}
}
