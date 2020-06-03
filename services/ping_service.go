package services

import "fmt"

type pingServiceInterface interface {
	// In order for any type to be pingServiceInterface it needs to implement
	// the following functions
	PingService() (string, error)
}

type pingServiceStruct struct{}

var (
	// PingServiceVar makes possible to use the method from other packages
	// Create a public variable of the interface type
	// and the value is going to be a new instance of the struct implementation
	PingServiceVar pingServiceInterface = pingServiceStruct{}
)

// PingService service is the actual service
func (service pingServiceStruct) PingService() (string, error) {
	// What if here we were calling a database or an external 3rd party?
	// do you imagine calling that every time you do unit testing?
	// In that case this thing might fail or not, so we need to return at least an error as well
	// so the controller can check the error
	fmt.Println("Doing some complex things ...")
	// as you can see we never return an error ...
	// so I have no way in the controller to test the problematic if statement
	return "pong", nil
}
