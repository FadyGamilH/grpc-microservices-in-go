gen-order:
	@protoc --proto_path=protobuf "protobuf/orders.proto" --go_out=services/common/protos/orders --go_opt=paths=source_relative --go-grpc_out=services/common/protos/orders --go-grpc_opt=paths=source_relative


run-orders:
	@go run services/orders/*.go