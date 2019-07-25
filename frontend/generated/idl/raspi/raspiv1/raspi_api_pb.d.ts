import * as jspb from "google-protobuf"

import * as protoc$gen$swagger_options_annotations_pb from '../../../protoc-gen-swagger/options/annotations_pb';
import * as google_api_annotations_pb from '../../../google/api/annotations_pb';
import * as gogo_pb from '../../../gogo_pb';
import * as idl_raspi_raspiv1_types_pb from '../../../idl/raspi/raspiv1/types_pb';

export class TempStreamRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TempStreamRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TempStreamRequest): TempStreamRequest.AsObject;
  static serializeBinaryToWriter(message: TempStreamRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TempStreamRequest;
  static deserializeBinaryFromReader(message: TempStreamRequest, reader: jspb.BinaryReader): TempStreamRequest;
}

export namespace TempStreamRequest {
  export type AsObject = {
  }
}

export class TempStreamResponse extends jspb.Message {
  getTemp(): idl_raspi_raspiv1_types_pb.Temperature | undefined;
  setTemp(value?: idl_raspi_raspiv1_types_pb.Temperature): void;
  hasTemp(): boolean;
  clearTemp(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TempStreamResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TempStreamResponse): TempStreamResponse.AsObject;
  static serializeBinaryToWriter(message: TempStreamResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TempStreamResponse;
  static deserializeBinaryFromReader(message: TempStreamResponse, reader: jspb.BinaryReader): TempStreamResponse;
}

export namespace TempStreamResponse {
  export type AsObject = {
    temp?: idl_raspi_raspiv1_types_pb.Temperature.AsObject,
  }
}

export class RadiatorRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RadiatorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RadiatorRequest): RadiatorRequest.AsObject;
  static serializeBinaryToWriter(message: RadiatorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RadiatorRequest;
  static deserializeBinaryFromReader(message: RadiatorRequest, reader: jspb.BinaryReader): RadiatorRequest;
}

export namespace RadiatorRequest {
  export type AsObject = {
  }
}

export class RadiatorResponse extends jspb.Message {
  getRadi(): idl_raspi_raspiv1_types_pb.Radiator | undefined;
  setRadi(value?: idl_raspi_raspiv1_types_pb.Radiator): void;
  hasRadi(): boolean;
  clearRadi(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RadiatorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RadiatorResponse): RadiatorResponse.AsObject;
  static serializeBinaryToWriter(message: RadiatorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RadiatorResponse;
  static deserializeBinaryFromReader(message: RadiatorResponse, reader: jspb.BinaryReader): RadiatorResponse;
}

export namespace RadiatorResponse {
  export type AsObject = {
    radi?: idl_raspi_raspiv1_types_pb.Radiator.AsObject,
  }
}

export class SetRadiatorRequest extends jspb.Message {
  getRadi(): idl_raspi_raspiv1_types_pb.Radiator | undefined;
  setRadi(value?: idl_raspi_raspiv1_types_pb.Radiator): void;
  hasRadi(): boolean;
  clearRadi(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetRadiatorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SetRadiatorRequest): SetRadiatorRequest.AsObject;
  static serializeBinaryToWriter(message: SetRadiatorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetRadiatorRequest;
  static deserializeBinaryFromReader(message: SetRadiatorRequest, reader: jspb.BinaryReader): SetRadiatorRequest;
}

export namespace SetRadiatorRequest {
  export type AsObject = {
    radi?: idl_raspi_raspiv1_types_pb.Radiator.AsObject,
  }
}

export class SetRadiatorResponse extends jspb.Message {
  getRadi(): idl_raspi_raspiv1_types_pb.Radiator | undefined;
  setRadi(value?: idl_raspi_raspiv1_types_pb.Radiator): void;
  hasRadi(): boolean;
  clearRadi(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetRadiatorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SetRadiatorResponse): SetRadiatorResponse.AsObject;
  static serializeBinaryToWriter(message: SetRadiatorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetRadiatorResponse;
  static deserializeBinaryFromReader(message: SetRadiatorResponse, reader: jspb.BinaryReader): SetRadiatorResponse;
}

export namespace SetRadiatorResponse {
  export type AsObject = {
    radi?: idl_raspi_raspiv1_types_pb.Radiator.AsObject,
  }
}

export class StatusRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StatusRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StatusRequest): StatusRequest.AsObject;
  static serializeBinaryToWriter(message: StatusRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StatusRequest;
  static deserializeBinaryFromReader(message: StatusRequest, reader: jspb.BinaryReader): StatusRequest;
}

export namespace StatusRequest {
  export type AsObject = {
  }
}

export class StatusResponse extends jspb.Message {
  getStatus(): idl_raspi_raspiv1_types_pb.Status | undefined;
  setStatus(value?: idl_raspi_raspiv1_types_pb.Status): void;
  hasStatus(): boolean;
  clearStatus(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StatusResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StatusResponse): StatusResponse.AsObject;
  static serializeBinaryToWriter(message: StatusResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StatusResponse;
  static deserializeBinaryFromReader(message: StatusResponse, reader: jspb.BinaryReader): StatusResponse;
}

export namespace StatusResponse {
  export type AsObject = {
    status?: idl_raspi_raspiv1_types_pb.Status.AsObject,
  }
}

