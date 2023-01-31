generate:
	# buf mod update
	# buf build
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	buf generate

build-pet:
	docker build -t pet -f app/pet/Dockerfile .

lint:
	buf lint
	buf breaking --against 'https://github.com/hungtran150/api-app.git#branch=master'

BUF_VERSION:=0.51.1

# Error TODO: fix this
install:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

deploy:
	docker compose up --scale order=2
