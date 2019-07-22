package echo_test

import (
	"app/api/echo/mocks"
	. "app/generated/idl/echo"
	"app/pkg/testconn"
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

		mockCtrl      *gomock.Controller
		mockAPIServer *mocks.MockEchoAPIServer

		client EchoAPIClient
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		_ = cancel

		mockCtrl = gomock.NewController(GinkgoT())
		mockAPIServer = mocks.NewMockEchoAPIServer(mockCtrl)

		buf := testconn.NewBufNet()

		register := func(s *grpc.Server) {
			RegisterEchoAPIServer(s, mockAPIServer)
		}
		go testconn.StartGRPCTestServer(ctx, buf, register)

		newClient := func(c *grpc.ClientConn) interface{} {
			return NewEchoAPIClient(c)
		}
		client = testconn.NewGRPCTestClient(ctx, buf, newClient).(EchoAPIClient)
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
					).Return(&EchoResponse{}, nil)
			})

			It("returns test reply", func() {
				reply, err := client.Echo(context.TODO(), &EchoRequest{})
				Expect(reply).To(BeAssignableToTypeOf(&EchoResponse{}))
				Expect(err).To(BeNil())
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
				reply, err := client.Echo(context.TODO(), &EchoRequest{})
				Expect(reply).To(BeNil())
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
