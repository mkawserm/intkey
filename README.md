# IntKey
Integer Key Example GRPC Service

# INSTALL 

```
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc

```

# GRPC GATEWAY

```
https://github.com/grpc-ecosystem/grpc-gateway

$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

```

# LOAD TEST #

```
GHZ: https://github.com/bojand/ghz
/intkey.IntKeyRPC/SafeIncrement

ghz --insecure --async \
  --proto ./pkg/proto/intkey/intkey.proto \
  --call intkey.IntKeyRPC/SafeIncrement \
  -c 1 -n 1 \
  -d '{"key":"test", "value": 1}' 0.0.0.0:3030 

```