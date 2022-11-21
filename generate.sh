
# Greet
protoc --proto_path=./greet/greetpb/  --go_out=./greet --go-grpc_out=./greet ./greet/greetpb/greet.proto

#Calculator
protoc --proto_path=./calculator/calculatorpb/  --go_out=./calculator --go-grpc_out=./calculator ./calculator/calculatorpb/calculator.proto

