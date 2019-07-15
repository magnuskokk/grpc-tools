package heartbeat

//go: generate mockgen -destination=mocks/service.go -package=mocks servicetool/services/heartbeat HeartbeatServiceServer

/*
	# Heartbeat service gRPC service Typescript client.
	#@protoc \
#		-I/usr/local/include \
#		-I. \
#	  	-I${GOPATH}/src \
 #   	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I./protobuf \
#		--js_out=import_style=commonjs:./services \
#		--grpc-web_out=import_style=typescript,mode=grpcwebtext:./services \
#		commander/service.proto
*/
