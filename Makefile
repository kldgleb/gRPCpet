create:
	protoc 	--go_out=pkg \
          	--go-grpc_out=pkg \
          	api/proto/*.proto

