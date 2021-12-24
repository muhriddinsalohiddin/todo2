package service

import (
	"log"
	"os"
	"testing"

	pb "github.com/muhriddinsalohiddin/todo2/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.TaskServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Did not connect grpc client", err)
	}
	client = pb.NewTaskServiceClient(conn)

	os.Exit(m.Run())
}
