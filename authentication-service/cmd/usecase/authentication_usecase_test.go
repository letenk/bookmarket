package usecase

import (
	"authentication_service/cmd/config"
	"authentication_service/cmd/models/domain"
	"authentication_service/cmd/models/web"
	"authentication_service/cmd/repository"
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/rizkydarmawan-letenk/jabufaker"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// truncateUsers as truncate table users
func truncateUsers(db *sql.DB) {
	db.Exec("TRUNCATE users")
}

// Function register random user
func RegisterRandomUser(t *testing.T) domain.User {
	// Load file .env
	godotenv.Load("../../.env")
	// Open connection
	// Connection to DB
	conn := config.SetupDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	defer conn.Close()

	// Truncate table users before test running
	truncateUsers(conn)

	// Used repository
	authenticationRepository := repository.NewRepositoryUser(conn)
	// Use usecase
	authenticationUseCase := NewUseCaseUser(authenticationRepository)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Getting random province and regencies for city
	province := jabufaker.RandomProvince()
	regency := jabufaker.RandomRegency(province)

	// Create sample data
	user := web.RegisterInput{
		Fullname: jabufaker.RandomPerson(),
		Email:    jabufaker.RandomEmail(),
		Address:  jabufaker.RandomString(20),
		City:     regency,
		Province: province,
		Mobile:   "082233377729",
		Password: "password",
		Role:     "admin",
	}

	// Register user
	newUser, err := authenticationUseCase.Register(ctx, user)
	if err != nil {
		log.Panic(err)
	}

	// Test pass
	assert.NotEmpty(t, newUser.ID)
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

// TestRegisterSuccess as test register random user is exist
func TestRegisterSuccess(t *testing.T) {
	RegisterRandomUser(t)
}

func TestRegisterFailEmailIsAvailable(t *testing.T) {
	// Register Random user
	newUser := RegisterRandomUser(t)

	// Load file .env
	godotenv.Load("../../.env")
	// Open connection
	// Connection to DB
	conn := config.SetupDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	defer conn.Close()

	// Used repository
	authenticationRepository := repository.NewRepositoryUser(conn)
	// Use usecase
	authenticationUseCase := NewUseCaseUser(authenticationRepository)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	user := web.RegisterInput{
		Fullname: newUser.Fullname,
		Email:    newUser.Email,
		Address:  newUser.Address,
		City:     newUser.City,
		Province: newUser.Province,
		Mobile:   newUser.Mobile,
		Password: newUser.Password,
		Role:     newUser.Role,
	}

	_, err := authenticationUseCase.Register(ctx, user)
	// Test pass
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "email already exist")
}
