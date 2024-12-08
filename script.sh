

gen() {
    echo "Generating protobuf files..."
    protoc \
        --go_out=protobuf \
        --go-grpc_out=protobuf --go-grpc_opt=require_unimplemented_servers=false --grpc-gateway_out=protobuf \
        protobuf/*.proto
}


# protoc --go_out=protobuf \
#  --go-grpc_out=protobuf --go-grpc_opt=require_unimplemented_servers=false\
#  --grpc-gateway_out=protobuf \
#  protobuf/brand.proto

protoc --go_out=protobuf \
 --go-grpc_out=protobuf --go-grpc_opt=require_unimplemented_servers=false\
 --grpc-gateway_out=protobuf \
 protobuf/brand.proto