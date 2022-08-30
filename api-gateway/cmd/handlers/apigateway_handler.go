package handlers

import (
	"api_gateway/cmd/models/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiGateway(c *gin.Context) {
	response := web.JSONResponseWithoutData(
		http.StatusOK,
		"success",
		"Api gateway is healthy.",
	)

	c.JSON(http.StatusOK, response)
}
