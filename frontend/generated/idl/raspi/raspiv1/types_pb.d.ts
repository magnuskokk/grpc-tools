import * as jspb from "google-protobuf"

import * as gogo_pb from '../../../gogo_pb';

export class Temperature extends jspb.Message {
  getReading(): number;
  setReading(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Temperature.AsObject;
  static toObject(includeInstance: boolean, msg: Temperature): Temperature.AsObject;
  static serializeBinaryToWriter(message: Temperature, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Temperature;
  static deserializeBinaryFromReader(message: Temperature, reader: jspb.BinaryReader): Temperature;
}

export namespace Temperature {
  export type AsObject = {
    reading: number,
  }
}

export class Radiator extends jspb.Message {
  getEnabled(): boolean;
  setEnabled(value: boolean): void;

  getLevel(): number;
  setLevel(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Radiator.AsObject;
  static toObject(includeInstance: boolean, msg: Radiator): Radiator.AsObject;
  static serializeBinaryToWriter(message: Radiator, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Radiator;
  static deserializeBinaryFromReader(message: Radiator, reader: jspb.BinaryReader): Radiator;
}

export namespace Radiator {
  export type AsObject = {
    enabled: boolean,
    level: number,
  }
}

export class Status extends jspb.Message {
  getTemperature(): Temperature | undefined;
  setTemperature(value?: Temperature): void;
  hasTemperature(): boolean;
  clearTemperature(): void;

  getRadiator(): Radiator | undefined;
  setRadiator(value?: Radiator): void;
  hasRadiator(): boolean;
  clearRadiator(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Status.AsObject;
  static toObject(includeInstance: boolean, msg: Status): Status.AsObject;
  static serializeBinaryToWriter(message: Status, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Status;
  static deserializeBinaryFromReader(message: Status, reader: jspb.BinaryReader): Status;
}

export namespace Status {
  export type AsObject = {
    temperature?: Temperature.AsObject,
    radiator?: Radiator.AsObject,
  }
}

