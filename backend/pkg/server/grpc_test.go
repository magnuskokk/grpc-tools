package server_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"

	"grpc-tools/pkg/server"
	"grpc-tools/pkg/testing"

	"grpc-tools/pkg/server/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
)

var _ = Describe("gRPC server and client for heartbeat service", func() {
	var (
		ctx    context.Context
		cancel context.CancelFunc

		mockCtrl          *gomock.Controller
		mockServiceServer *mocks.MockHeartbeatServiceServer

		client   server.HeartbeatServiceClient
		testPong *server.PingReply
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		_ = cancel

		mockCtrl = gomock.NewController(GinkgoT())
		mockServiceServer = mocks.NewMockHeartbeatServiceServer(mockCtrl)

		buf := testing.NewBufNet()

		register := func(s *grpc.Server) {
			server.RegisterHeartbeatServiceServer(s, mockServiceServer)
		}
		go testing.StartGRPCTestServer(ctx, buf, register)

		newClient := func(c *grpc.ClientConn) interface{} {
			return server.NewHeartbeatServiceClient(c)
		}
		client = testing.NewGRPCTestClient(ctx, buf, newClient).(server.HeartbeatServiceClient)

		testPong = &server.PingReply{
			Message: []byte("test"),
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
					Ping(
						gomock.Any(),
						gomock.AssignableToTypeOf(&server.PingRequest{})).
					Return(testPong, nil)
			})

			It("returns test reply", func() {
				reply, err := client.Ping(context.TODO(), &server.PingRequest{})
				Expect(reply.Message).To(Equal([]byte("test")))
				Expect(err).To(BeNil())
			})
		})

		Context("Sending fails", func() {
			BeforeEach(func() {
				mockServiceServer.EXPECT().
					Ping(
						gomock.Any(),
						gomock.AssignableToTypeOf(&server.PingRequest{})).
					Return(nil, errors.New("service error"))
			})

			It("returns error", func() {
				reply, err := client.Ping(context.TODO(), &server.PingRequest{})
				Expect(reply).To(BeAssignableToTypeOf(&server.PingReply{}))
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
