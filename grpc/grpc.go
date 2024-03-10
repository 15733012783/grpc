package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

func RegisterGRPC(ip string, port int, c func(s *grpc.Server)) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", ip, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	c(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
