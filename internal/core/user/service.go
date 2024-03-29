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
	"strings"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

type Service struct {
	randInstance *mathRand.Rand
}

func NewService() *Service {
	randInstance := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))

	return &Service{
		randInstance: randInstance,
	}
}

func (s *Service) CreateUser(_ context.Context, u domain.User) (string, error) {
	u.Username = s.createUsername(u.FirstName, u.LastName, u.InstitutionID, u.ListNumber)
	u.Password = s.createPassword(u.Password)
	u.DateJoined = time.Now()

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

func (s *Service) UpdateUser(_ context.Context, u domain.User) error {
	return nil
}

func (s *Service) DeleteUser(_ context.Context, id uint64) error {
	return nil
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
