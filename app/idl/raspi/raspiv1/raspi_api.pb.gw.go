// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: idl/raspi/raspiv1/raspi_api.proto

/*
Package raspiv1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package raspiv1

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_RaspiAPI_TempStream_0(ctx context.Context, marshaler runtime.Marshaler, client RaspiAPIClient, req *http.Request, pathParams map[string]string) (RaspiAPI_TempStreamClient, runtime.ServerMetadata, error) {
	var protoReq TempStreamRequest
	var metadata runtime.ServerMetadata

	stream, err := client.TempStream(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

func request_RaspiAPI_Radiator_0(ctx context.Context, marshaler runtime.Marshaler, client RaspiAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RadiatorRequest
	var metadata runtime.ServerMetadata

	msg, err := client.Radiator(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_RaspiAPI_SetRadiator_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_RaspiAPI_SetRadiator_0(ctx context.Context, marshaler runtime.Marshaler, client RaspiAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SetRadiatorRequest
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_RaspiAPI_SetRadiator_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.SetRadiator(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_RaspiAPI_Status_0(ctx context.Context, marshaler runtime.Marshaler, client RaspiAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StatusRequest
	var metadata runtime.ServerMetadata

	msg, err := client.Status(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterRaspiAPIHandlerFromEndpoint is same as RegisterRaspiAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterRaspiAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterRaspiAPIHandler(ctx, mux, conn)
}

// RegisterRaspiAPIHandler registers the http handlers for service RaspiAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterRaspiAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterRaspiAPIHandlerClient(ctx, mux, NewRaspiAPIClient(conn))
}

// RegisterRaspiAPIHandlerClient registers the http handlers for service RaspiAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "RaspiAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "RaspiAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "RaspiAPIClient" to call the correct interceptors.
func RegisterRaspiAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client RaspiAPIClient) error {

	mux.Handle("GET", pattern_RaspiAPI_TempStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_RaspiAPI_TempStream_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RaspiAPI_TempStream_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_RaspiAPI_Radiator_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_RaspiAPI_Radiator_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RaspiAPI_Radiator_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_RaspiAPI_SetRadiator_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_RaspiAPI_SetRadiator_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RaspiAPI_SetRadiator_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_RaspiAPI_Status_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_RaspiAPI_Status_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RaspiAPI_Status_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_RaspiAPI_TempStream_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"temp"}, ""))

	pattern_RaspiAPI_Radiator_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"radiator"}, ""))

	pattern_RaspiAPI_SetRadiator_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"radiator"}, ""))

	pattern_RaspiAPI_Status_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"status"}, ""))
)

var (
	forward_RaspiAPI_TempStream_0 = runtime.ForwardResponseStream

	forward_RaspiAPI_Radiator_0 = runtime.ForwardResponseMessage

	forward_RaspiAPI_SetRadiator_0 = runtime.ForwardResponseMessage

	forward_RaspiAPI_Status_0 = runtime.ForwardResponseMessage
)