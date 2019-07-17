package heartbeat_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"

	"app/pkg/testconn"
	"app/services/heartbeat"
	"app/services/heartbeat/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
)

var _ = Describe("gRPC server and client for heartbeat service", func() {
	var (
		ctx    context.Context
		cancel context.CancelFunc

		mockCtrl          *gomock.Controller
		mockServiceServer *mocks.MockServiceServer

		client   heartbeat.ServiceClient
		testPong *heartbeat.PingReply
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		_ = cancel

		mockCtrl = gomock.NewController(GinkgoT())
		mockServiceServer = mocks.NewMockServiceServer(mockCtrl)

		buf := testconn.NewBufNet()

		register := func(s *grpc.Server) {
			heartbeat.RegisterServiceServer(s, mockServiceServer)
		}
		go testconn.StartGRPCTestServer(ctx, buf, register)

		newClient := func(c *grpc.ClientConn) interface{} {
			return heartbeat.NewServiceClient(c)
		}
		client = testconn.NewGRPCTestClient(ctx, buf, newClient).(heartbeat.ServiceClient)

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
