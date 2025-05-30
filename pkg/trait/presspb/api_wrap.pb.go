// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package presspb

import (
	traits "github.com/smart-core-os/sc-api/go/traits"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapApi	adapts a traits.PressApiServer	and presents it as a traits.PressApiClient
func WrapApi(server traits.PressApiServer) *ApiWrapper {
	conn := wrap.ServerToClient(traits.PressApi_ServiceDesc, server)
	client := traits.NewPressApiClient(conn)
	return &ApiWrapper{
		PressApiClient: client,
		server:         server,
		conn:           conn,
		desc:           traits.PressApi_ServiceDesc,
	}
}

type ApiWrapper struct {
	traits.PressApiClient

	server traits.PressApiServer
	conn   grpc.ClientConnInterface
	desc   grpc.ServiceDesc
}

// UnwrapServer returns the underlying server instance.
func (w *ApiWrapper) UnwrapServer() traits.PressApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *ApiWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *ApiWrapper) UnwrapService() (grpc.ClientConnInterface, grpc.ServiceDesc) {
	return w.conn, w.desc
}
