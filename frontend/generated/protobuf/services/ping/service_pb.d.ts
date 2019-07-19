import * as jspb from "google-protobuf"

import * as google_api_annotations_pb from '../../../google/api/annotations_pb';
import * as protoc$gen$swagger_options_annotations_pb from '../../../protoc-gen-swagger/options/annotations_pb';
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from '../../../github.com/gogo/protobuf/gogoproto/gogo_pb';

export class PingRequest extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PingRequest): PingRequest.AsObject;
  static serializeBinaryToWriter(message: PingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingRequest;
  static deserializeBinaryFromReader(message: PingRequest, reader: jspb.BinaryReader): PingRequest;
}

export namespace PingRequest {
  export type AsObject = {
    message: string,
  }
}

export class PingReply extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingReply.AsObject;
  static toObject(includeInstance: boolean, msg: PingReply): PingReply.AsObject;
  static serializeBinaryToWriter(message: PingReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingReply;
  static deserializeBinaryFromReader(message: PingReply, reader: jspb.BinaryReader): PingReply;
}

export namespace PingReply {
  export type AsObject = {
    message: string,
  }
}

export class SomeRequest extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getCount(): number;
  setCount(value: number): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SomeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SomeRequest): SomeRequest.AsObject;
  static serializeBinaryToWriter(message: SomeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SomeRequest;
  static deserializeBinaryFromReader(message: SomeRequest, reader: jspb.BinaryReader): SomeRequest;
}

export namespace SomeRequest {
  export type AsObject = {
    id: Uint8Array | string,
    count: number,
    message: string,
  }
}

export class SomeReply extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SomeReply.AsObject;
  static toObject(includeInstance: boolean, msg: SomeReply): SomeReply.AsObject;
  static serializeBinaryToWriter(message: SomeReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SomeReply;
  static deserializeBinaryFromReader(message: SomeReply, reader: jspb.BinaryReader): SomeReply;
}

export namespace SomeReply {
  export type AsObject = {
    id: Uint8Array | string,
  }
}

