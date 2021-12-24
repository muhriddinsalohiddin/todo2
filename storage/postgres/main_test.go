package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/muhriddinsalohiddin/todo2/config"
	"github.com/muhriddinsalohiddin/todo2/pkg/db"
	"github.com/muhriddinsalohiddin/todo2/pkg/logger"
)

var pgRepo *taskRepo

func TestMain(m *testing.M) {
	cfg := config.Load()
	conn, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connect error", logger.Error(err))
	}
	pgRepo = NewTaskRepo(conn)
	os.Exit(m.Run())
}
