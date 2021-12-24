package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/muhriddinsalohiddin/todo2/config"
	pb "github.com/muhriddinsalohiddin/todo2/genproto"
	"github.com/muhriddinsalohiddin/todo2/pkg/db"
	"github.com/muhriddinsalohiddin/todo2/pkg/logger"
	"github.com/muhriddinsalohiddin/todo2/service"
	"github.com/muhriddinsalohiddin/todo2/storage"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "template-service")

	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres err", logger.Error(err))
	}

	pgStorage := storage.NewStoragePg(connDB)

	taskService := service.NewTaskService(pgStorage, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Connection GRPC error", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, taskService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening", logger.Error(err))
	}
}
