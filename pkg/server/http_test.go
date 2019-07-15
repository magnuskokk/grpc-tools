package heartbeat_test

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"

	"grpc-tools/pkg/testing"

	"grpc-tools/services/heartbeat"
	"grpc-tools/services/heartbeat/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
)

var _ = Describe("HTTP server and client for heartbeat service", func() {
	var (
		ctx    context.Context
		cancel context.CancelFunc

		mockCtrl          *gomock.Controller
		mockServiceServer *mocks.MockHeartbeatServiceServer

		server    *httptest.Server
		testReply *heartbeat.PingReply
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		_ = cancel

		mockCtrl = gomock.NewController(GinkgoT())
		mockServiceServer = mocks.NewMockHeartbeatServiceServer(mockCtrl)

		buf := testing.NewBufNet()

		go func() {
			register := func(s *grpc.Server) {
				heartbeat.RegisterHeartbeatServiceServer(s, mockServiceServer)
			}
			if err := testing.StartGRPCTestServer(ctx, buf, register); err != nil {
				log.Fatal(err)
			}
		}()

		var err error
		server, err = testing.NewGatewayTestServer(ctx, buf, heartbeat.RegisterHeartbeatServiceHandlerFromEndpoint)
		Expect(err).To(BeNil())

		testReply = &heartbeat.PingReply{
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
					Return(testReply, nil)
			})

			It("returns test reply", func() {
				res, err := http.Get(server.URL + "/ping")
				Expect(err).To(BeNil())

				Expect(res.StatusCode).To(Equal(http.StatusOK))

				body, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				Expect(err).To(BeNil())

				var reply map[string]interface{}
				err = json.Unmarshal(body, &reply)
				Expect(err).To(BeNil())

				// dGVzdA== is base64 encoded value for []byte("test") in json
				// we would have to unmarshal into a proper struct with []byte type
				// not use a map to auto convert back to bytes
				Expect(reply).To(HaveKeyWithValue("Message", "dGVzdA=="))
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
				res, err := http.Get(server.URL + "/ping")
				Expect(err).To(BeNil())

				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})
