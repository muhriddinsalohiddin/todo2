package service

import (
	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/muhriddinsalohiddin/todo2/genproto"
	log "github.com/muhriddinsalohiddin/todo2/pkg/logger"
	"github.com/muhriddinsalohiddin/todo2/storage"
)

type TaskService struct {
	storage storage.IStorage
	logger  log.Logger
}

func NewTaskService(storage storage.IStorage, log log.Logger) *TaskService {
	return &TaskService{
		storage: storage,
		logger:  log,
	}
}

func (t *TaskService) Create(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	id, err := uuid.NewV4()
	if err != nil {
		t.logger.Error("failed to generate uuid", log.Error(err))
		return nil, status.Error(codes.Internal, "failed to generate uuid")
	}
	req.Id = id.String()

	task, err := t.storage.Task().Create(*req)
	if err != nil {
		t.logger.Error("failed to create task", log.Error(err))
		return nil, status.Error(codes.Internal, "failed to create task")
	}
	return &task, nil
}

func (t *TaskService) Get(ctx context.Context, in *pb.ByIdReq) (*pb.Task, error) {
	task, err := t.storage.Task().Get(in.Id)
	if err != nil {
		t.logger.Error("failed to get task", log.Error(err))
		return nil, status.Error(codes.Internal, "failed to get task")
	}
	return &task, nil
}

func (t *TaskService) List(ctx context.Context, in *pb.ListReq) (*pb.ListResp, error) {
	list, err := t.storage.Task().List(*in)
	if err != nil {
		t.logger.Error("failed to get list task", log.Error(err))
		return nil, status.Error(codes.Internal, "failed to get list task")
	}
	return &list, nil
}

func (t *TaskService) Update(ctx context.Context, in *pb.Task) (*pb.Task, error) {
	task, err := t.storage.Task().Update(*in)
	if err != nil {
		t.logger.Error("failed to update task", log.Error(err))
		return nil, status.Error(codes.Internal, "failed to update task")
	}
	return &task, nil
}

func (t *TaskService) Delete(ctx context.Context, in *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := t.storage.Task().Delete(in.Id)
	if err != nil {
		t.logger.Error("failed to delete task", log.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete task")
	}
	return &pb.EmptyResp{}, nil
}

func (t *TaskService) ListOverdue(ctx context.Context, in *pb.ListOverReq) (*pb.ListOverResp, error) {
	list, err := t.storage.Task().ListOverdue(*in)
	if err != nil {
		t.logger.Error("failed to get listoverdue", log.Error(err))
		return &pb.ListOverResp{}, status.Error(codes.Internal, "failed to get listoverdue")
	}
	return &list, nil
}
