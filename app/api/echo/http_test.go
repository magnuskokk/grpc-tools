package echo_test

import (
	"app/api/echo/mocks"
	. "app/generated/idl/echo"
	"app/pkg/testconn"
	context "context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	grpc "google.golang.org/grpc"
)

var _ = Describe("HTTP server and client for echo service", func() {
	var (
		ctx    context.Context
		cancel context.CancelFunc

		mockCtrl      *gomock.Controller
		mockAPIServer *mocks.MockEchoAPIServer

		testServer   *httptest.Server
		testResponse *EchoResponse
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		_ = cancel

		mockCtrl = gomock.NewController(GinkgoT())
		mockAPIServer = mocks.NewMockEchoAPIServer(mockCtrl)

		buf := testconn.NewBufNet()

		go func() {
			register := func(s *grpc.Server) {
				RegisterEchoAPIServer(s, mockAPIServer)
			}
			if err := testconn.StartGRPCTestServer(ctx, buf, register); err != nil {
				log.Fatal(err)
			}
		}()

		var err error
		testServer, err = testconn.NewGatewayTestServer(ctx, buf, RegisterEchoAPIHandlerFromEndpoint)
		Expect(err).To(BeNil())

		testResponse = &EchoResponse{
			Message: "test",
		}
	})

	JustAfterEach(func() {
		mockCtrl.Finish()
		cancel()
	})

	Describe("Sending commands", func() {
		Context("Sending succeeds", func() {
			BeforeEach(func() {
				mockAPIServer.EXPECT().
					Echo(
						gomock.Any(),
						gomock.AssignableToTypeOf(&EchoRequest{}),
					).Return(testResponse, nil)
			})

			It("returns test reply", func() {
				res, err := http.Get(testServer.URL + "/echo")
				Expect(err).To(BeNil())

				Expect(res.StatusCode).To(Equal(http.StatusOK))

				body, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				Expect(err).To(BeNil())

				var reply map[string]interface{}
				err = json.Unmarshal(body, &reply)
				Expect(err).To(BeNil())

				Expect(reply).To(HaveKeyWithValue("message", "test"))
			})
		})

		Context("Sending fails", func() {
			BeforeEach(func() {
				mockAPIServer.EXPECT().
					Echo(
						gomock.Any(),
						gomock.AssignableToTypeOf(&EchoRequest{}),
					).Return(nil, errors.New("service error"))
			})

			It("returns error", func() {
				res, err := http.Get(testServer.URL + "/echo")
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				Expect(err).To(BeNil())
			})
		})
	})
})
