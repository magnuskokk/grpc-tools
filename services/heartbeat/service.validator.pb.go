// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/heartbeat/service.proto

package heartbeat

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/gogo/googleapis/google/api"
	time "time"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

func (this *PingRequest) Validate() error {
	return nil
}
func (this *PingReply) Validate() error {
	return nil
}
func (this *StreamRequest) Validate() error {
	if !(len(this.ID) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`Must have len(id) between 0 and 1000`))
	}
	if !(len(this.ID) < 1000) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`Must have len(id) between 0 and 1000`))
	}
	if this.CreateDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreateDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreateDate", err)
		}
	}
	return nil
}
func (this *StreamPacket) Validate() error {
	return nil
}
