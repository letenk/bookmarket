package tests

import (
	"authentication_service/cmd/config"
	"authentication_service/cmd/models/domain"
	"authentication_service/cmd/repository"
	"authentication_service/pkg"
	"context"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/rizkydarmawan-letenk/jabufaker"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// InsertRandomUser as function insert random user and if success return new user
func InsertRandomUser(t *testing.T) domain.User {
	// Load file .env
	godotenv.Load("../.env")
	// Open connection
	// Connection to DB
	conn := config.SetupDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	defer conn.Close()

	// Truncate table users before test running
	pkg.TruncateUsers(conn)

	// Used repository
	repo := repository.NewRepositoryUser(conn)

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Generate uuid
	id := uuid.New()
	pass := "password"
	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		log.Panicf("Hashing password failed, err: %s\n", err)
	}

	// Getting random province and regencies for city
	province := jabufaker.RandomProvince()
	regency := jabufaker.RandomRegency(province)

	// Create sample data
	user := domain.User{
		ID:       id,
		Fullname: jabufaker.RandomPerson(),
		Email:    jabufaker.RandomEmail(),
		Address:  jabufaker.RandomString(20),
		City:     regency,
		Province: province,
		Mobile:   "082233377728",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	// Insert user
	newUser, err := repo.Insert(ctx, user)
	if err != nil {
		log.Panic(err)
	}

	// Test pass
	assert.Equal(t, user.ID, newUser.ID)
	assert.Equal(t, user.Fullname, newUser.Fullname)
	assert.Equal(t, user.Email, newUser.Email)
	assert.Equal(t, user.Address, newUser.Address)
	assert.Equal(t, user.City, newUser.City)
	assert.Equal(t, user.Province, newUser.Province)
	assert.Equal(t, user.Mobile, newUser.Mobile)
	assert.Equal(t, user.Role, newUser.Role)
	assert.NotEmpty(t, newUser.CreatedAt)
	assert.NotEmpty(t, newUser.UpdatedAt)

	// Test pass password, but before test compare first
	err = bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte("password"))
	assert.Nil(t, err)

	return newUser
}

// TestInsertUserSuccess as testing function InsertRandomUser is success
func TestInsertUserSuccess(t *testing.T) {
	InsertRandomUser(t)
}

// TestCheckEmailIsAvailable as check email is available if yes, return true
func TestCheckEmailIsAvailable(t *testing.T) {
	// Load file .env
	godotenv.Load("../.env")
	// Open connection
	// Connection to DB
	conn := config.SetupDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	defer conn.Close()

	// Truncate table users before test running
	pkg.TruncateUsers(conn)

	// Used repository
	repo := repository.NewRepositoryUser(conn)

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Insert Random user
	newUser := InsertRandomUser(t)

	// Find by email
	checkEmail := repo.EmailIsAvailable(ctx, newUser.Email)

	// Test pass
	assert.True(t, checkEmail)
}

// TestCheckEmailIsNotAvailable as check email is available if yes, return false
func TestCheckEmailIsNotAvailable(t *testing.T) {
	// Load file .env
	godotenv.Load("../.env")
	// Open connection
	// Connection to DB
	conn := config.SetupDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	defer conn.Close()

	// Truncate table users before test running
	pkg.TruncateUsers(conn)

	// Used repository
	repo := repository.NewRepositoryUser(conn)

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Find by email
	checkEmail := repo.EmailIsAvailable(ctx, "fail@test.com")

	// Test Pass
	assert.False(t, checkEmail)
}