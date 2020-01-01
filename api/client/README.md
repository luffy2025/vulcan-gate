# Vulcan Frame

**Vulcan Frame** is a production-ready game server framework designed for building high-performance, highly available game server clusters using microservices architecture.

# Vulcan API Client

**vulcan-api-client** defines the API protocol for game clients using Protocol Buffer IDL, following [Kratos API Definition](https://go-kratos.dev/docs/component/api/) standards.

## Features

- Protocol-first design using Protobuf
- Modular architecture for game services
- High-performance networking layer
- Distributed system ready
- Client compatibility package

## Vulcan API Client

**vulcan-api-client** defines the API protocol for game clients using Protocol Buffer IDL, following [Kratos API Definition](https://go-kratos.dev/docs/component/api/) standards.

## Installation

```bash
Install protocol buffer compiler
brew install protobuf
Install buf tool for protocol management
brew install bufbuild/buf/buf
```

## Protocol Structure

### Module Definition

Modules define business domains using Protobuf enums:

```protobuf
enum ModuleID {
  ModuleUnknown = 0;
  Room = 1;
}
```

### Sequence Definition

Sequences define message IDs within modules:

```protobuf
// sequence/user.proto
enum RoomSeq {
	RoomUnknown = 0;
	CreateRoom = 1;
}
```

### Message Definition

Messages follow standard gRPC service definitions:

```protobuf
// message/room.proto
service RoomService {
rpc CreateRoom(CreateRoomRequest) returns (CreateRoomResponse);
rpc JoinRoom(JoinRoomRequest) returns (JoinRoomResponse);
}
```

### Packet Structure

Network packet format definition:

```protobuf
// packet/packet.proto
message Packet {
  // ...
  int32 mod = 4; // Module ID, globally unique
  int32 seq = 5; // Message ID within the module, unique within the module
  int64 obj = 6; // Module object ID, according to the business agreement to pass the corresponding object ID
  bytes data = 7; // Serialized bytes of the cs/sc protocol in the message
  // ...
}
```

## Code Generation

vulcan-api-client will be introduced as a submodule in projects like vulcan-gate and vulcan-game, and the projects will provide tools to generate the corresponding code.

## Contributing

If you have any suggestions or feedback, please feel free to open an issue or submit a `pull request`.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
