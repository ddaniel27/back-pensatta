package postgres

import (
	"context"
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

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	var inst models.InstitutionModel
	if err := r.db.WithContext(ctx).Where("code = ?", u.InstitutionCode).First(&inst).Error; err != nil {
		return err
	}

	um := toUserModel(u)
	um.InstitutionID = inst.ID

	return r.db.WithContext(ctx).Create(&um).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id uint64) (domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Get(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetByRole(ctx context.Context, role string) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.WithContext(ctx).Where("role = ?", role).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Raw(`
 SELECT u.*, COALESCE(l.value, 'es') AS language
 FROM pensatta_user u
 LEFT JOIN pensatta_languages l ON u.institution_id = l.institution_id
 WHERE u.username = $1`, username).
		Scan(&user).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetProfileResumen(ctx context.Context, user domain.User) (map[string]interface{}, error) {
	var totalRecords int64
	var institution models.InstitutionModel
	var calification models.CalificationModel

	if err := r.db.WithContext(ctx).
		Model(&models.RecordModel{}).
		Where("user_id = ?", user.ID).
		Count(&totalRecords).Error; err != nil {
		return nil, err
	}

	if totalRecords == 0 {
		totalRecords = 1
	}

	if err := r.db.WithContext(ctx).
		Where("id = ?", user.InstitutionID).
		First(&institution).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		Where("user_id = ?", user.ID).
		First(&calification).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}

		calification = models.CalificationModel{
			UserID:       user.ID,
			AverageScore: 0,
			AverageTime:  0,
		}
	}

	return map[string]interface{}{
		"institution_name": institution.Name,
		"average_score":    calification.AverageScore,
		"average_time":     calification.AverageTime,
		"total_exercises":  totalRecords,
	}, nil
}

func (r *UserRepository) Update(ctx context.Context, u domain.User) error {
	var um models.UserModel
	r.db.WithContext(ctx).Table("pensatta_user").Where("username = ?", u.Username).First(&um)

	myUM := toUserModel(u)
	myUM.ID = um.ID
	return r.db.WithContext(ctx).Save(&myUM).Error
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&models.UserModel{}, id).Error
}

func toUserModel(u domain.User) models.UserModel {
	return models.UserModel{
		Username:      u.Username,
		Password:      u.Password,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		InstitutionID: u.InstitutionID,
		Role:          u.Role,
		ListNumber:    u.ListNumber,
		DateJoined:    u.DateJoined,
		LastLogin:     u.LastLogin,
	}
}
