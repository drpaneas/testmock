package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPingController(t *testing.T) {
	// the PingController(c *gin.Context) receives *gin.Context
	// so we will test it by creating a fake one via gin.CreateTestContext
	// this one needs an http.ResponseWriter as a input
	// so we have to create the ResponseWrite first
	// to do this we will use the httptest.NewRecorder()
	fakeResponseWriter := httptest.NewRecorder()
	fakeGinContext, _ := gin.CreateTestContext(fakeResponseWriter)

	// Now that we have all the we need, we can actually call the function
	// by passing the fake context to test it
	PingController(fakeGinContext)

	// The first test
	// The PingController() returns a message based on the HTTP return code
	// so let's test this behavior

	if fakeResponseWriter.Code != http.StatusOK {
		t.Error("response code should be 200")
	}

	// The Second test
	// If the above test was passed, it means it returned 200
	// So now we need to check if it will print the message as it should be
	if fakeResponseWriter.Body.String() != "pong" {
		t.Error("response body string should say 'pong")
	}
}
