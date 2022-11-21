
# Getting started
protoc --proto_path=./greet/greetpb/  --go_out=./greet --go-grpc_out=./greet ./greet/greetpb/greet.proto

