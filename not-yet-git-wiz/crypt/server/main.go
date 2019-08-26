package server

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/wizardsoftheweb/git-wiz/not-yet-git-wiz/crypt/proto"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Auth(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Printf("Received: %v", in.GetKey())
	return &pb.AuthResponse{Key: "Hello " + in.GetKey()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAgentServer(s, &server{})
	go func() {
		time.Sleep(30 * time.Second)
		log.Printf("Exiting")
		s.GracefulStop()
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
