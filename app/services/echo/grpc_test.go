package echo_test

import (
	"app/pkg/testconn"
	"app/services/echo"
	"app/services/echo/mocks"
	context "context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	grpc "google.golang.org/grpc"
)

var _ = Describe("gRPC server and client for echo service", func() {
	var (
		ctx    context.Context
		cancel context.CancelFunc

		mockCtrl          *gomock.Controller
		mockServiceServer *mocks.MockEchoServiceServer

		client      echo.EchoServiceClient
		testRequest *echo.EchoRequest
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		_ = cancel

		mockCtrl = gomock.NewController(GinkgoT())
		mockServiceServer = mocks.NewMockEchoServiceServer(mockCtrl)

		buf := testconn.NewBufNet()

		register := func(s *grpc.Server) {
			echo.RegisterEchoServiceServer(s, mockServiceServer)
		}
		go testconn.StartGRPCTestServer(ctx, buf, register)

		newClient := func(c *grpc.ClientConn) interface{} {
			return echo.NewEchoServiceClient(c)
		}
		client = testconn.NewGRPCTestClient(ctx, buf, newClient).(echo.EchoServiceClient)

		testRequest = &echo.EchoRequest{
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
				mockServiceServer.EXPECT().
					Echo(
						gomock.Any(),
						gomock.AssignableToTypeOf(&echo.EchoRequest{}),
					).Return(testRequest, nil)
			})

			It("returns test reply", func() {
				reply, err := client.Echo(context.TODO(), testRequest)
				Expect(*reply).To(Equal(*testRequest))
				Expect(err).To(BeNil())
			})
		})

		Context("Sending fails", func() {
			BeforeEach(func() {
				mockServiceServer.EXPECT().
					Echo(
						gomock.Any(),
						gomock.AssignableToTypeOf(&echo.EchoRequest{}),
					).Return(nil, errors.New("service error"))
			})

			It("returns error", func() {
				reply, err := client.Echo(context.TODO(), testRequest)
				Expect(reply).To(BeNil())
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
