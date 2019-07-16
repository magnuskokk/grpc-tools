/**
 * @fileoverview gRPC-Web generated client stub for server
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as google_api_annotations_pb from '../../../google/api/annotations_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as protoc$gen$swagger_options_annotations_pb from '../../../protoc-gen-swagger/options/annotations_pb';
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from '../../../github.com/gogo/protobuf/gogoproto/gogo_pb';

import {
  PingReply,
  PingRequest,
  StreamPacket,
  StreamRequest} from './service_pb';

export class HeartbeatServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoPing = new grpcWeb.AbstractClientBase.MethodInfo(
    PingReply,
    (request: PingRequest) => {
      return request.serializeBinary();
    },
    PingReply.deserializeBinary
  );

  ping(
    request: PingRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: PingReply) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/server.HeartbeatService/Ping',
      request,
      metadata || {},
      this.methodInfoPing,
      callback);
  }

  methodInfoStream = new grpcWeb.AbstractClientBase.MethodInfo(
    StreamPacket,
    (request: StreamRequest) => {
      return request.serializeBinary();
    },
    StreamPacket.deserializeBinary
  );

  stream(
    request: StreamRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/server.HeartbeatService/Stream',
      request,
      metadata || {},
      this.methodInfoStream);
  }

}

