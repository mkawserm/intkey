VERSION = "0.0.1"
change-version:
	@echo $(VERSION)>VERSION
	@echo "package constant\n\n//Version constant of the service\nconst Version = \"$(VERSION)\"">pkg/constant/version.go

test:
	go test -race ./... -v

update-module:
	go get -v google.golang.org/grpc
	go get -v github.com/rs/zerolog
	go get -v github.com/caarlos0/env/v6
	go get -v google.golang.org/protobuf
#	go env -w GOPRIVATE=repo
#	go get -v repo

build:
#	GOPRIVATE=repo GIT_TERMINAL_PROMPT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/intkey cmd/intkey/intkey.go
	go build -v -o bin/intkey cmd/intkey/intkey.go

simple-build:
	go build -v -o bin/intkey cmd/intkey/intkey.go

build-with-proxy:
#	GOPRIVATE=repo GIT_TERMINAL_PROMPT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://proxy.golang.org go build -v -o bin/intkey cmd/intkey/intkey.go

clean:
	@rm -rf bin/intkey
	@rm -rf .proto-dir

run:
	go run cmd/intkey/intkey.go

alpine-container:
	docker build . -f docker/alpine/Dockerfile -t intkey:latest --build-arg GIT_DOMAIN=$GIT_DOMAIN --build-arg GIT_USERNAME=$GIT_USERNAME --build-arg GIT_PASSWORD=$GIT_PASSWORD

slim-container:
	docker build . -f docker/slim/Dockerfile -t intkey:latest-slim --build-arg GIT_DOMAIN=$GIT_DOMAIN --build-arg GIT_USERNAME=$GIT_USERNAME --build-arg GIT_PASSWORD=$GIT_PASSWORD


protoc:
	@protoc \
		-I=./pkg/proto \
		--go_opt=module=github.com/mkawserm/intkey \
		--go_out=. \
		--go-grpc_opt=module=github.com/mkawserm/intkey \
		--go-grpc_out=. \
		./pkg/proto/intkey/intkey.proto

push:
	git push github master
