build:
	@echo "Build file"
protoc:
	@echo "Build protobuf file"
	@protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=paths=source_relative,plugins=grpc:. \
		--grpc-gateway_out=paths=source_relative,logtostderr=true:. \
		--swagger_out=:. \
		delivery/grpc/user/user.proto
