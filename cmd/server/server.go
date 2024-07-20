package main

import (
	"google.golang.org/grpc"
	"github.com/KubeOperator/kobe/api"
	"github.com/KubeOperator/kobe/pkg/server"
	"net"
)

func newTcpListener(address string) (*net.Listener, error) {
	s, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
func newServer() *grpc.Server {
	options := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(100 * 1024 * 1024 * 1024),
		grpc.MaxSendMsgSize(100 * 1024 * 1024 * 1024),
	}
	gs := grpc.NewServer(options...)
	kobe := server.NewKobe()
	api.RegisterKobeApiServer(gs, kobe)
	return gs
}
