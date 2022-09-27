package usecase

import (
	"authentication_service/cmd/models/domain"
	"authentication_service/cmd/models/web"
	"authentication_service/cmd/repository"
	"context"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	Register(input web.RegisterInput) (domain.User, error)
	Login(input web.LoginInput) (string, error)
}

type useCase struct {
	repository repository.Repository
}

func NewUseCaseUser(repository repository.Repository) *useCase {
	return &useCase{repository}
}

func (u *useCase) Register(input web.RegisterInput) (domain.User, error) {
	// create context with timeout duration 3 second
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
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

type Claim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func (u *useCase) Login(input web.LoginInput) (string, error) {
	// Create context with timeout duration 3 second
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Find user by email
	user, err := u.repository.FindByEmail(ctx, input.Email)
	if user.ID == uuid.Nil {
		return "", errors.New("email or password incorrect")
	}

	if err != nil {
		return "", err
	}

	// If user available, compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("email or password incorrect")
	}

	// If the password is matched, Generate token jwt
	// Create 1 day token active
	expirationTime := time.Now().AddDate(0, 0, 1)

	// Create claim for payload tokne
	claim := Claim{
		UserID: user.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// Get secret key
	SecretJWT := os.Getenv("SECRET_JWT")
	// Sifned token with secret key
	signedToken, err := token.SignedString([]byte(SecretJWT))
	if err != nil {
		return "", err
	}

	// If success, return token
	return signedToken, nil
}
