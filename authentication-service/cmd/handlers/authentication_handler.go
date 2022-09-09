package handlers

import (
	"authentication_service/cmd/models/web"
	"authentication_service/cmd/usecase"
	"authentication_service/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	useCase usecase.UseCase
}

func NewHandler(useCase usecase.UseCase) *handler {
	return &handler{useCase}
}

func (h *handler) Register(c *gin.Context) {
	var input web.RegisterInput

	// Get Payload
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := pkg.ValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Register failed.",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Register
	_, err = h.useCase.Register(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Register failed.",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create format response
	response := web.ApiResponseWithoutData(
		http.StatusCreated,
		"success",
		"You have successfully registered",
	)

	c.JSON(http.StatusCreated, response)
}
