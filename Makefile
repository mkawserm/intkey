VERSION = "0.0.1"
change-version:
	@echo $(VERSION)>VERSION
	@echo "package constant\n\n//Version constant of the service\nconst Version = \"$(VERSION)\"">pkg/constant/version.go

update-module:
	go get -v github.com/golang/protobuf
	go get -v google.golang.org/grpc
	go get -v github.com/rs/zerolog
	go get -v github.com/caarlos0/env/v6
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

#container:
#	docker build . -f Dockerfile -t intkey:latest --build-arg GITLAB_ID=$GITLAB_ID --build-arg GITLAB_TOKEN=$GITLAB_TOKEN

#slim-container:
#	docker build . -f Dockerfile:slim -t intkey:latest-slim --build-arg GITLAB_ID=$GITLAB_ID --build-arg GITLAB_TOKEN=$GITLAB_TOKEN


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
