package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/muhriddinsalohiddin/todo2/storage/postgres"
	"github.com/muhriddinsalohiddin/todo2/storage/repo"
)

// IStorage ...
type IStorage interface {
	Task() repo.TaskStorageI
}

type storagePg struct {
	db       *sqlx.DB
	taskRepo repo.TaskStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		taskRepo: postgres.NewTaskRepo(db),
	}
}

func (s storagePg) Task() repo.TaskStorageI {
	return s.taskRepo
}
