module github.com/Sri2103/services/internal/cart

go 1.22.1

replace github.com/Sri2103/services/pkg => ../../pkg/

require (
	github.com/Sri2103/services/pkg v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.62.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)
