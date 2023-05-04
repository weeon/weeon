package grpc

import (
	"net"

	"google.golang.org/grpc"
)

type ServerInterface interface {
	Serve(lis net.Listener) error
	grpc.ServiceRegistrar
}
