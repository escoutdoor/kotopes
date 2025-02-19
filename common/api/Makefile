LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.20.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.20.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:		
	make generate_auth
	make generate_user
	make generate_access
	make generate_pet
	make generate_favorite;

generate_auth:
	protoc --proto_path auth/v1 --proto_path vendor.protogen \
	--go_out=auth/v1 --go_opt=paths=source_relative \
	--go-grpc_out=auth/v1 --go-grpc_opt=paths=source_relative \
	--validate_out lang=go:auth/v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	auth/v1/auth.proto

generate_user:
	protoc --proto_path user/v1 --proto_path vendor.protogen \
	--go_out=user/v1 --go_opt=paths=source_relative \
	--go-grpc_out=user/v1 --go-grpc_opt=paths=source_relative \
	--validate_out lang=go:user/v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	user/v1/user.proto

generate_access:
	protoc --proto_path access/v1 --proto_path vendor.protogen \
	--go_out=access/v1 --go_opt=paths=source_relative \
	--go-grpc_out=access/v1 --go-grpc_opt=paths=source_relative \
	--validate_out lang=go:access/v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	access/v1/access.proto

generate_pet:
	protoc --proto_path pet/v1 --proto_path vendor.protogen \
	--go_out=pet/v1 --go_opt=paths=source_relative \
	--go-grpc_out=pet/v1 --go-grpc_opt=paths=source_relative \
	--validate_out lang=go:pet/v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	pet/v1/pet.proto

generate_favorite:
	protoc --proto_path favorite/v1 --proto_path pet/v1 --proto_path vendor.protogen \
	--go_out=favorite/v1 --go_opt=paths=source_relative \
	--go-grpc_out=favorite/v1 --go-grpc_opt=paths=source_relative \
	--validate_out lang=go:favorite/v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	favorite/v1/favorite.proto

vendor-proto:
	@if [ ! -d vendor.protogen/protovalidate ]; then \
		git clone -b main --single-branch --depth=1 --filter=tree:0 \
			https://github.com/bufbuild/protovalidate vendor.protogen/protovalidate &&\
		mv vendor.protogen/protovalidate/proto/protovalidate/buf vendor.protogen &&\
		rm -rf vendor.protogen/protovalidate ;\
	fi
	@if [ ! -d vendor.protogen/google ]; then \
		git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
			https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
		mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
		git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
		mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
		rm -rf vendor.protogen/openapiv2 ;\
	fi
