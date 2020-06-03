package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/testmock/services"
)

// Here we are going to implement our fake method
type mockPingServiceStruct struct {
	// create a function sig so we can fill it later
	FakePingService func() (string, error)
}

func (mock mockPingServiceStruct) PingService() (string, error) {
	fmt.Println("I am faking the complex things.")
	return mock.FakePingService()
}

func TestPingControllerNoError(t *testing.T) {

	// the PingController(c *gin.Context) receives *gin.Context
	// so we will test it by creating a fake one via gin.CreateTestContext
	// this one needs an http.ResponseWriter as a input
	// so we have to create the ResponseWrite first
	// to do this we will use the httptest.NewRecorder()
	fakeResponseWriter := httptest.NewRecorder()
	fakeGinContext, _ := gin.CreateTestContext(fakeResponseWriter)

	// Now that we have all the we need, we can actually call the function
	// by passing the fake context to test it. That is the PingController
	// which calls it like this: services.PingServiceVar.PingService()
	// HACK services.PingServiceVar to point to the fake Service

	// Create a local variable and inster the structure
	mockPingController := mockPingServiceStruct{}
	mockPingController.FakePingService = func() (string, error) {
		return "pong", nil
	}
	services.PingServiceVar = mockPingController // <-- this is the hack

	// Now I can mess around with the function
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

func TestPingControllerWithError(t *testing.T) {
	// Copy paste everything from the previous function

	fakeResponseWriter := httptest.NewRecorder()
	fakeGinContext, _ := gin.CreateTestContext(fakeResponseWriter)

	mockPingController := mockPingServiceStruct{}
	mockPingController.FakePingService = func() (string, error) {
		err := fmt.Errorf(http.StatusText(http.StatusInternalServerError))
		return "", err
	}
	services.PingServiceVar = mockPingController // <-- this is the hack
	PingController(fakeGinContext)

	// The first test
	// The PingController() returns a message based on the HTTP return code
	// so let's test this behavior if there is an error

	if fakeResponseWriter.Code != http.StatusInternalServerError {
		t.Error("response code should not be 200")
	}

	// The Second test
	// If the above test was passed, it means it didn't return 200
	// So now it should not print the message
	if fakeResponseWriter.Body.String() == "pong" {
		t.Error("response body string should NOT say 'pong")
	}
}
