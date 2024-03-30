package repositories

import "pensatta/internal/core/domain"

type UserRepository interface {
	Create(u domain.User) error
	GetByID(id uint64) (domain.User, error)
	Get() ([]domain.User, error)
	GetByRole(role string) ([]domain.User, error)
	Update(u domain.User) error
	Delete(id uint64) error
}
