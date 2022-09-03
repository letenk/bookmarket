package repository

import (
	"authentication_service/cmd/config"
	"authentication_service/cmd/models/domain"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestInsertUserSuccess(t *testing.T) {
	// Open connection
	// Connection to DB
	conn := config.SetupDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	defer conn.Close()

	// Used repository
	repo := NewRepositoryUser(conn)

	// Generate uuid
	id := uuid.New()
	pass := "password"
	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		log.Panicf("Hashing password failed, err: %s\n", err)
	}

	// Create sample data
	user := domain.User{
		ID:       id,
		Fullname: "Rizky Darmawan",
		Email:    "letenk@test.com",
		Address:  "Jl. Jalan",
		City:     "Binjai",
		Province: "Sumatera Utara",
		Mobile:   "082277760694",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	// Insert user
	newUser, err := repo.Insert(user)
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

}
