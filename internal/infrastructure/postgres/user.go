package postgres

import (
	"pensatta/internal/core/domain"
	"pensatta/internal/infrastructure/postgres/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u domain.User) error {
	var inst models.InstitutionModel
	if err := r.db.Where("code = ?", u.InstitutionCode).First(&inst).Error; err != nil {
		return err
	}

	um := toUserModel(u)
	um.InstitutionID = inst.ID

	return r.db.Create(&um).Error
}

func (r *UserRepository) GetByID(id uint64) (domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Get() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetByRole(role string) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Where("role = ?", role).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Update(u domain.User) error {
	return r.db.Save(toUserModel(u)).Error
}

func (r *UserRepository) Delete(id uint64) error {
	return r.db.Delete(&models.UserModel{}, id).Error
}

func toUserModel(u domain.User) models.UserModel {
	return models.UserModel{
		Username:   u.Username,
		Password:   u.Password,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Role:       u.Role,
		ListNumber: u.ListNumber,
		DateJoined: u.DateJoined,
		LastLogin:  u.LastLogin,
	}
}
