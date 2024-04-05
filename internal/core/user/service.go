package user

import (
	"context"
	cryptoRand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math"
	mathRand "math/rand"
	"pensatta/internal/core/domain"
	"pensatta/internal/core/ports/repositories"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

type Service struct {
	randInstance *mathRand.Rand
	userRepo     repositories.UserRepository
}

func NewService(ur repositories.UserRepository) *Service {
	randInstance := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))

	return &Service{
		randInstance: randInstance,
		userRepo:     ur,
	}
}

func (s *Service) CreateUser(ctx context.Context, u domain.User) (string, error) {
	u.Username = s.createUsername(u.FirstName, u.LastName, u.InstitutionCode, u.ListNumber)
	u.Password = s.createPassword(u.Password)
	u.DateJoined = time.Now()

	if err := s.userRepo.Create(ctx, u); err != nil {
		return "", err
	}

	return u.Username, nil
}

func (s *Service) GetUser(_ context.Context, id uint64) (domain.User, error) {
	return domain.User{}, nil
}

func (s *Service) GetUsers(_ context.Context) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (s *Service) GetUserByRole(_ context.Context, role string) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (s *Service) GetUserProfileResumen(ctx context.Context, user domain.User) (map[string]interface{}, error) {
	res, err := s.userRepo.GetProfileResumen(ctx, user)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"institution_name": res["institution_name"],
		"last_login":       user.LastLogin,
		"resumen": map[string]interface{}{
			"total_exercises": res["total_exercises"],
			"average_score":   res["average_score"],
			"average_time":    res["average_time"],
		},
	}, nil
}

func (s *Service) UpdateUser(_ context.Context, u domain.User) error {
	return nil
}

func (s *Service) DeleteUser(_ context.Context, id uint64) error {
	return nil
}

func (s *Service) ValidateCredentials(ctx context.Context, username, password string) (domain.User, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return domain.User{}, err
	}

	if !s.validatePassword(password, user.Password) {
		return domain.User{}, fmt.Errorf("invalid credentials")
	}

	user.LastLogin = time.Now()
	if err := s.userRepo.Update(ctx, user); err != nil {
		return domain.User{}, err
	}

	user.Password = ""
	return user, nil
}

func (s *Service) validatePassword(password, encryptedPassword string) bool {
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 4 {
		return false
	}

	passLength := 32
	iterations, salt, hash := parts[1], parts[2], parts[3]
	iter, err := strconv.Atoi(iterations)
	if err != nil {
		return false
	}

	hashBytes := pbkdf2.Key([]byte(password), []byte(salt), iter, passLength, sha256.New)
	newHash := base64.StdEncoding.EncodeToString(hashBytes)

	return newHash == hash
}

func (s *Service) createUsername(firstname, lastname, institutionCode string, listNumber uint64) string {
	randomNumber := s.randInstance.Float64()
	randomNumber = math.Floor(randomNumber*(999-100) + 100)

	username := fmt.Sprintf("%s%.0f%s%s%d", institutionCode, randomNumber, firstname[:2], lastname[:2], listNumber)
	return strings.ToUpper(username)
}

func (s *Service) createPassword(password string) string {
	saltBytes := make([]byte, 16)
	if _, err := cryptoRand.Read(saltBytes); err != nil {
		fmt.Println(err)
	}

	salt := base64.StdEncoding.EncodeToString(saltBytes)
	iterations := 39000
	passLength := 32
	algorithm := "sha256"
	encrypt := "pbkdf2"

	hashBytes := pbkdf2.Key([]byte(password), []byte(salt), iterations, passLength, sha256.New)
	hash := base64.StdEncoding.EncodeToString(hashBytes)

	return fmt.Sprintf("%s_%s$%d$%s$%s", encrypt, algorithm, iterations, salt, hash)
}
