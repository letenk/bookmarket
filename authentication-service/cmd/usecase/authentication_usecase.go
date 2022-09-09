package usecase

import (
	"authentication_service/cmd/models/domain"
	"authentication_service/cmd/models/web"
	"authentication_service/cmd/repository"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	Register(ctx context.Context, input web.RegisterInput) (domain.User, error)
}

type useCase struct {
	repository repository.Repository
}

func NewUseCaseUser(repository repository.Repository) *useCase {
	return &useCase{repository}
}

func (u *useCase) Register(ctx context.Context, input web.RegisterInput) (domain.User, error) {
	// create context with timeout duration 3 second
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	// Check email is available
	checkEmail := u.repository.EmailIsAvailable(ctx, input.Email)
	// If checkEmail is true, return error
	if checkEmail {
		return domain.User{}, errors.New("email already exist")
	}

	// Pass data request into domain user
	user := domain.User{}
	user.Fullname = input.Fullname
	user.Email = input.Email
	user.Address = input.Address
	user.City = input.City
	user.Province = input.Province
	user.Mobile = input.Mobile
	user.Role = input.Role

	//  Generate uuid
	id := uuid.New()
	user.ID = id

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	// passing password hash into object domain user
	user.Password = string(passwordHash)

	// Insert new user
	newUser, err := u.repository.Insert(ctx, user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
