package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/testmock/services"
)

// PingController checks the return code from HTTP status
// it returns the result of the PingService if it's HTTP 200
// it returns nothing if it's HTTP 404
func PingController(c *gin.Context) {
	result, err := services.PingService()
	if err != nil {
		// So now we cannot test this, unless we fake the service to Fail on purpose
		// we need to have complete function of what services.PingService() returns
		// but we cannot do this, because Go is a compiled language and typed defined
		// you cannot mock this function
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.String(http.StatusOK, result)
	}
}
