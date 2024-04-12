gen:
  @protoc \
    --proto_path=protobuf "protobuf/orders.proto" \
    --go_out=services/common/protos --go_opt=paths=source_relative \
    --go-grpc_out=services/common/protos 
    --go-grpc_opt=paths=source_relative