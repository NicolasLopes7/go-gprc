protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.
