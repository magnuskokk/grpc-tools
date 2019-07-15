package heartbeat_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"

	"grpc-tools/pkg/testing"

	"grpc-tools/services/heartbeat"
	"grpc-tools/services/heartbeat/mocks"

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

		client   heartbeat.HeartbeatServiceClient
		testPong *heartbeat.PingReply
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		_ = cancel

		mockCtrl = gomock.NewController(GinkgoT())
		mockServiceServer = mocks.NewMockHeartbeatServiceServer(mockCtrl)

		buf := testing.NewBufNet()

		register := func(s *grpc.Server) {
			heartbeat.RegisterHeartbeatServiceServer(s, mockServiceServer)
		}
		go testing.StartGRPCTestServer(ctx, buf, register)

		newClient := func(c *grpc.ClientConn) interface{} {
			return heartbeat.NewHeartbeatServiceClient(c)
		}
		client = testing.NewGRPCTestClient(ctx, buf, newClient).(heartbeat.HeartbeatServiceClient)

		testPong = &heartbeat.PingReply{
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
						gomock.AssignableToTypeOf(&heartbeat.PingRequest{})).
					Return(testPong, nil)
			})

			It("returns test reply", func() {
				reply, err := client.Ping(context.TODO(), &heartbeat.PingRequest{})
				Expect(reply.Message).To(Equal([]byte("test")))
				Expect(err).To(BeNil())
			})
		})

		Context("Sending fails", func() {
			BeforeEach(func() {
				mockServiceServer.EXPECT().
					Ping(
						gomock.Any(),
						gomock.AssignableToTypeOf(&heartbeat.PingRequest{})).
					Return(nil, errors.New("service error"))
			})

			It("returns error", func() {
				reply, err := client.Ping(context.TODO(), &heartbeat.PingRequest{})
				Expect(reply).To(BeAssignableToTypeOf(&heartbeat.PingReply{}))
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
