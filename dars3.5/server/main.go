package main

import (
	"context"
	pb "n11/Firdavs/dars3.5/genproto/generator"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGeneratorServer
}
func (s *server) RandomPicker(ctx context.Context, in *pb.Response)(*pb)  {
	
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := server{}
	grpc := grpc.NewServer()
	pb.RegisterGeneratorServer(grpc, &s)

	err = grpc.Serve(listener)
	if err != nil {
		panic(err)
	}
}
