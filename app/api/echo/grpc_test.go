package echo_test

import (
	"app/api/echo/mocks"
	. "app/idl/echo/echov1"
	"app/pkg/server"
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
		go server.StartGRPCServer(ctx, buf.Listener, register)

		newClient := func(c *grpc.ClientConn) interface{} {
			return NewEchoAPIClient(c)
		}
		cl, err := server.NewGRPCClient(server.ClientOptions{
			Ctx:           ctx,
			Addr:          "bufnet",
			Dialer:        buf.DialContext,
			ClientFactory: newClient,
		})
		Expect(err).To(BeNil())

		var ok bool
		client, ok = cl.(EchoAPIClient)
		Expect(ok).To(BeTrue())
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
					).Return(&EchoResponse{Message: "test"}, nil)
			})

			It("returns test reply", func() {
				reply, err := client.Echo(context.TODO(), &EchoRequest{})
				Expect(err).To(BeNil())
				Expect(reply.GetMessage()).To(Equal("test"))
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
