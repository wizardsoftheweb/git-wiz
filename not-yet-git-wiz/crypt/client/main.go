package client

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "github.com/wizardsoftheweb/git-wiz/not-yet-git-wiz/crypt/proto"
)

const (
	address    = "localhost:50051"
	defaultKey = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAgentClient(conn)

	// Contact the server and print out its response.
	key := defaultKey
	if len(os.Args) > 1 {
		key = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Auth(ctx, &pb.AuthRequest{Key: key})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Key: %s", r.GetKey())
}
