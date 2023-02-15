# dockertool-legend

## Simple tool which is aimed to maintain statistic over working docker containers and send it in stream to clients

### Features:
- gRPC server Streaming
- Graceful shutdown
- Hex architecture pattern

## How to start up project:
- Start server with default parameters:
``make server``
- Start client with 15 seconds timeout:
``make client``

## Flags:
### Server:
- [console = *bool*] *stream stats to StdOut* default to false
- [port = *int*] *gRPC port to listen* default to 50051
### Client:
- [timeout = *int*] *timeout to RPC call* default to 0
- [port = *int*] *gRPC server port* default to 50051