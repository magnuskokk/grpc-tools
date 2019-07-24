/**
 * @fileoverview gRPC-Web generated client stub for app.echo.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as protoc$gen$swagger_options_annotations_pb from '../../../protoc-gen-swagger/options/annotations_pb';
import * as google_api_annotations_pb from '../../../google/api/annotations_pb';
import * as gogo_pb from '../../../gogo_pb';

import {
  EchoRequest,
  EchoResponse} from './echo_api_pb';

export class EchoAPIClient {
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

  methodInfoEcho = new grpcWeb.AbstractClientBase.MethodInfo(
    EchoResponse,
    (request: EchoRequest) => {
      return request.serializeBinary();
    },
    EchoResponse.deserializeBinary
  );

  echo(
    request: EchoRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: EchoResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/app.echo.v1.EchoAPI/Echo',
      request,
      metadata || {},
      this.methodInfoEcho,
      callback);
  }

}

