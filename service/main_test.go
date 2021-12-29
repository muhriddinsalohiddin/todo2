package service

import (
	"log"
	"os"
	"testing"

	"github.com/muhriddinsalohiddin/todo2/config"
	pb "github.com/muhriddinsalohiddin/todo2/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.TaskServiceClient

func TestMain(m *testing.M) {
	cfg := config.Load()
	conn, err := grpc.Dial("localhost"+cfg.RPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Did not connect grpc client", err)
	}
	client = pb.NewTaskServiceClient(conn)

	os.Exit(m.Run())
}
