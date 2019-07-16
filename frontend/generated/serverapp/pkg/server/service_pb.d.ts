import * as jspb from "google-protobuf"

import * as google_api_annotations_pb from '../../../google/api/annotations_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as protoc$gen$swagger_options_annotations_pb from '../../../protoc-gen-swagger/options/annotations_pb';
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from '../../../github.com/gogo/protobuf/gogoproto/gogo_pb';

export class PingRequest extends jspb.Message {
  getMessage(): Uint8Array | string;
  getMessage_asU8(): Uint8Array;
  getMessage_asB64(): string;
  setMessage(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PingRequest): PingRequest.AsObject;
  static serializeBinaryToWriter(message: PingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingRequest;
  static deserializeBinaryFromReader(message: PingRequest, reader: jspb.BinaryReader): PingRequest;
}

export namespace PingRequest {
  export type AsObject = {
    message: Uint8Array | string,
  }
}

export class PingReply extends jspb.Message {
  getMessage(): Uint8Array | string;
  getMessage_asU8(): Uint8Array;
  getMessage_asB64(): string;
  setMessage(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingReply.AsObject;
  static toObject(includeInstance: boolean, msg: PingReply): PingReply.AsObject;
  static serializeBinaryToWriter(message: PingReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingReply;
  static deserializeBinaryFromReader(message: PingReply, reader: jspb.BinaryReader): PingReply;
}

export namespace PingReply {
  export type AsObject = {
    message: Uint8Array | string,
  }
}

export class StreamRequest extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getCreateDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreateDate(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasCreateDate(): boolean;
  clearCreateDate(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StreamRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StreamRequest): StreamRequest.AsObject;
  static serializeBinaryToWriter(message: StreamRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StreamRequest;
  static deserializeBinaryFromReader(message: StreamRequest, reader: jspb.BinaryReader): StreamRequest;
}

export namespace StreamRequest {
  export type AsObject = {
    id: Uint8Array | string,
    createDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class StreamPacket extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  getSum(): Uint8Array | string;
  getSum_asU8(): Uint8Array;
  getSum_asB64(): string;
  setSum(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StreamPacket.AsObject;
  static toObject(includeInstance: boolean, msg: StreamPacket): StreamPacket.AsObject;
  static serializeBinaryToWriter(message: StreamPacket, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StreamPacket;
  static deserializeBinaryFromReader(message: StreamPacket, reader: jspb.BinaryReader): StreamPacket;
}

export namespace StreamPacket {
  export type AsObject = {
    id: Uint8Array | string,
    data: Uint8Array | string,
    sum: Uint8Array | string,
  }
}

