package handlers

import (
	"authentication_service/cmd/models/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	response := web.JSONResponseWithoutData(
		http.StatusOK,
		"success",
		"Func register is ready.",
	)

	c.JSON(http.StatusOK, response)
}
