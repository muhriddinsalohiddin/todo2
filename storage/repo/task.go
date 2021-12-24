package repo

import (
	pb "github.com/muhriddinsalohiddin/todo2/genproto"
)

// TaskStorageI
type TaskStorageI interface {
	Create(pb.Task) (pb.Task, error)
	Get(id string) (pb.Task, error)
	List(pb.ListReq) (pb.ListResp, error)
	Update(pb.Task) (pb.Task, error)
	Delete(id string) error
	ListOverdue(pb.ListOverReq) (pb.ListOverResp, error)
}
