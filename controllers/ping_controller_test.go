package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mock_services "github.com/testmock/mocks"
	"github.com/testmock/services"
)

var _ = Describe("Testing the PingController", func() {
	var (
		// Define mock controller
		mockCtrl           *gomock.Controller
		mockPingController *mock_services.MockpingServiceInterface
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockPingController = mock_services.NewMockpingServiceInterface(mockCtrl)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Context("When hits /ping", func() {
		It("is expected to return pong and no error", func() {
			fakeResponseWriter := httptest.NewRecorder()
			fakeGinContext, _ := gin.CreateTestContext(fakeResponseWriter)
			mockPingController.EXPECT().PingService().Return("pong", nil)
			services.PingServiceVar = mockPingController
			PingController(fakeGinContext)
			Expect(fakeResponseWriter.Code).To(Equal(http.StatusOK))
			Expect(fakeResponseWriter.Body.String()).To(Equal("pong"))
		})
	})
	Context("When doesn't hit /ping", func() {
		It("is expected to return nothing and throw an error", func() {
			fakeResponseWriter := httptest.NewRecorder()
			fakeGinContext, _ := gin.CreateTestContext(fakeResponseWriter)
			err := fmt.Errorf(http.StatusText(http.StatusInternalServerError))
			mockPingController.EXPECT().PingService().Return("", err)
			services.PingServiceVar = mockPingController
			PingController(fakeGinContext)
			Expect(fakeResponseWriter.Code).ToNot(Equal(http.StatusOK))
			Expect(fakeResponseWriter.Body.String()).ToNot(Equal("pong"))
		})
	})
})
