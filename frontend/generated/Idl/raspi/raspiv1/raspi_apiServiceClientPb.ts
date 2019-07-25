/**
 * @fileoverview gRPC-Web generated client stub for app.raspi.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as protoc$gen$swagger_options_annotations_pb from '../../../protoc-gen-swagger/options/annotations_pb';
import * as google_api_annotations_pb from '../../../google/api/annotations_pb';
import * as gogo_pb from '../../../gogo_pb';
import * as idl_raspi_raspiv1_types_pb from '../../../idl/raspi/raspiv1/types_pb';

import {
  RadiatorRequest,
  RadiatorResponse,
  SetRadiatorRequest,
  SetRadiatorResponse,
  StatusRequest,
  StatusResponse,
  TempStreamRequest,
  TempStreamResponse} from './raspi_api_pb';

export class RaspiAPIClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; }) {
    if (!options) options = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoTempStream = new grpcWeb.AbstractClientBase.MethodInfo(
    TempStreamResponse,
    (request: TempStreamRequest) => {
      return request.serializeBinary();
    },
    TempStreamResponse.deserializeBinary
  );

  tempStream(
    request: TempStreamRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/app.raspi.v1.RaspiAPI/TempStream',
      request,
      metadata || {},
      this.methodInfoTempStream);
  }

  methodInfoRadiator = new grpcWeb.AbstractClientBase.MethodInfo(
    RadiatorResponse,
    (request: RadiatorRequest) => {
      return request.serializeBinary();
    },
    RadiatorResponse.deserializeBinary
  );

  radiator(
    request: RadiatorRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: RadiatorResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/app.raspi.v1.RaspiAPI/Radiator',
      request,
      metadata || {},
      this.methodInfoRadiator,
      callback);
  }

  methodInfoSetRadiator = new grpcWeb.AbstractClientBase.MethodInfo(
    SetRadiatorResponse,
    (request: SetRadiatorRequest) => {
      return request.serializeBinary();
    },
    SetRadiatorResponse.deserializeBinary
  );

  setRadiator(
    request: SetRadiatorRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: SetRadiatorResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/app.raspi.v1.RaspiAPI/SetRadiator',
      request,
      metadata || {},
      this.methodInfoSetRadiator,
      callback);
  }

  methodInfoStatus = new grpcWeb.AbstractClientBase.MethodInfo(
    StatusResponse,
    (request: StatusRequest) => {
      return request.serializeBinary();
    },
    StatusResponse.deserializeBinary
  );

  status(
    request: StatusRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: StatusResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/app.raspi.v1.RaspiAPI/Status',
      request,
      metadata || {},
      this.methodInfoStatus,
      callback);
  }

}

