package grpc

import (
	"context"
	"log"
	"net"
	pb "ricochetrobots/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	*grpc.Server
}

func New() *Server {
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterWalkingSkeletonServer(s, &WalkinSkeletonServer{})
	return &Server{s}
}

func (s *Server) Run(port string) {
	lis, err := net.Listen("tcp", "localhost:" + port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

type WalkinSkeletonServer struct {
	pb.UnimplementedWalkingSkeletonServer
}

func (s *WalkinSkeletonServer) Walk(context.Context, *pb.WalkRequest) (*pb.WalkReply, error) {
	return &pb.WalkReply{Message: "walk"}, nil
}
