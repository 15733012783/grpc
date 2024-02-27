package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RegisterGRPC(port int, c func(s *grpc.Server)) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	c(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
