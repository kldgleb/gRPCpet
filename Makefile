create:
	protoc 	--go_out=pkg \
          	--go-grpc_out=pkg \
          	api/proto/*.proto
migrate up:
	migrate -path ./db/migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up
migrate down:
	migrate -path ./db/migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' down
