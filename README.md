## gRPC Template - Golang

-   `apt install -y protobuf-compiler`
-   `protoc --version`

---

### Protocol compiler setup

-   `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
-   `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
-   `export PATH="$PATH:$(go env GOPATH)/bin"`

---

### Compile protobuffers

-   `export PATH="$PATH:$(go env GOPATH)/bin"`
-   `protoc --go_out=. --go-grpc_out=. proto/greet.proto`
-   `go mod tidy`
