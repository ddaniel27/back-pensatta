package repositories

import "pensatta/internal/core/domain"

type UserRepository interface {
	CreateUser(u domain.User) error
	GetUser(id uint64) (domain.User, error)
	GetUsers() ([]domain.User, error)
	GetUserByRole(role string) ([]domain.User, error)
	UpdateUser(u domain.User) error
	DeleteUser(id uint64) error
}
