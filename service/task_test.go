package service

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/muhriddinsalohiddin/todo2/genproto"
)

func TestTaskService_Create(t *testing.T) {
	tests := []struct {
		name  string
		input pb.Task
		want  pb.Task
	}{
		{
			name: "Testing Create tasks",
			input: pb.Task{
				Assignee: "Assignee",
				Title:    "Title",
				Summary:  "Summary",
				Deadline: "2022-04-04",
				Status:   "Aktive",
			},
			want: pb.Task{
				Assignee: "Assignee",
				Title:    "Title",
				Summary:  "Summary",
				Deadline: "2022-04-04T00:00:00Z",
				Status:   "Aktive",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := client.Create(context.Background(), &test.input)
			got.Id = ""
			got.CreatedAt = ""
			got.DeletedAt = ""
			got.UpdatedAt = ""
			if err != nil {
				t.Error("Failed create tasks", err)
			}
			if !reflect.DeepEqual(test.want, *got) {
				t.Fatalf("%s: Expected:%v, got:%s", test.name, test.want, got)
			}
		})
	}
}

func TestTaskService_Get(t *testing.T) {
	taskcheck := pb.Task{
		Assignee: "Assignee",
		Title:    "Get",
		Summary:  "Summary",
		Deadline: "2022-04-04T00:00:00Z",
		Status:   "Aktive",
	}
	task, _ := client.Create(context.Background(), &taskcheck)
	tests := []struct {
		name  string
		input pb.ByIdReq
		want  pb.Task
	}{
		{
			name: "get taks by id",
			input: pb.ByIdReq{
				Id: task.Id,
			},
			want: taskcheck,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := client.Get(context.Background(), &test.input)
			if err != nil {
				t.Error("Failed get task", err)
			}
			got.Id = ""
			got.UpdatedAt = ""
			got.CreatedAt = ""
			if !reflect.DeepEqual(test.want, *got) {
				t.Fatalf("%s: Expected:%v, got:%v", test.name, test.want, got)
			}
		})
	}
	client.Delete(context.Background(), &pb.ByIdReq{Id: task.Id})
}

func TestTaskService_Update(t *testing.T) {
	task, _ := client.Create(context.Background(), &pb.Task{
		Assignee: "Assignee",
		Title:    "Update",
		Summary:  "Summary",
		Deadline: "2022-04-04",
		Status:   "Aktive",
	})
	tests := []struct {
		name  string
		input pb.Task
		want  pb.Task
	}{
		{
			name: "update task",
			input: pb.Task{
				Id:       task.Id,
				Assignee: "Assignee1",
				Title:    "Update1",
				Summary:  "Summary1",
				Deadline: "2022-04-03",
				Status:   "Deaktive",
			},
			want: pb.Task{
				Id:       task.Id,
				Assignee: "Assignee1",
				Title:    "Update1",
				Summary:  "Summary1",
				Deadline: "2022-04-03",
				Status:   "Deaktive",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := client.Update(context.Background(), &test.input)
			if err != nil {
				t.Error("Failed update task", err)
			}
			if reflect.DeepEqual(test.want.UpdatedAt, got.UpdatedAt) {
				t.Fatalf("%s: Expected:%v, got:%v", test.name, test.want.UpdatedAt, got.UpdatedAt)
			}
		})
	}
	client.Delete(context.Background(), &pb.ByIdReq{Id: task.Id})
}

func TestTaskService_Delete(t *testing.T) {
	task, _ := client.Create(context.Background(), &pb.Task{
		Assignee: "Assignee",
		Title:    "Update",
		Summary:  "Summary",
		Deadline: "2022-04-04",
		Status:   "Aktive",
	})
	tests := []struct {
		name  string
		input pb.ByIdReq
		want  pb.EmptyResp
	}{
		{
			name: "Delete task",
			input: pb.ByIdReq{
				Id: task.Id,
			},
			want: pb.EmptyResp{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got, err := client.Delete(context.Background(), &test.input)
			if err != nil {
				t.Fatalf("failed delete task: %v", err)
			}
			if reflect.DeepEqual(got, test.want) {
				t.Fatalf("%s: Expected:%v, got:%v", test.name, test.want, got)
			}
		})
	}
}

func TestTaskService_ListOverdue(t *testing.T) {
	createTask1, _ := client.Create(context.Background(), &pb.Task{
		Assignee: "Assignee",
		Title:    "Title",
		Summary:  "Summary",
		Deadline: "2022-04-03",
		Status:   "Aktive",
	})
	createTask2, _ := client.Create(context.Background(), &pb.Task{
		Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e2",
		Assignee: "Assignee",
		Title:    "Title",
		Summary:  "Summary",
		Deadline: "2022-01-04",
		Status:   "Aktive",
	})

	tests := []struct {
		name  string
		input *pb.ListOverReq
		want  *pb.ListOverResp
	}{
		{
			name: "ListOverdue",
			input: &pb.ListOverReq{
				Time:  "2020-04-04",
				Page:  1,
				Limit: 10,
			},
			want: &pb.ListOverResp{
				Tasks: []*pb.Task{
					createTask1, createTask2,
				},
				Count: 2,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := client.ListOverdue(context.Background(), test.input)
			if err != nil {
				t.Fatalf("Failed list Overdue task err:%v", err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("%s: expected:%v, got:%v", test.name, test.want, got)
			}
		})
	}
	client.Delete(context.Background(), &pb.ByIdReq{Id: createTask1.Id})
	client.Delete(context.Background(), &pb.ByIdReq{Id: createTask2.Id})
}

func TestTaskService_List(t *testing.T) {
	createTask1, _ := client.Create(context.Background(), &pb.Task{
		Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e5",
		Assignee: "Assignee",
		Title:    "Title",
		Summary:  "Summary",
		Deadline: "2022-04-03",
		Status:   "Aktive",
	})
	createTask2, _ := client.Create(context.Background(), &pb.Task{
		Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e2",
		Assignee: "Assignee",
		Title:    "Title",
		Summary:  "Summary",
		Deadline: "2022-01-04",
		Status:   "Aktive",
	})

	tests := []struct {
		name  string
		input *pb.ListReq
		want  *pb.ListResp
	}{
		{
			name: "List",
			input: &pb.ListReq{
				Page:  1,
				Limit: 10,
			},
			want: &pb.ListResp{
				Tasks: []*pb.Task{
					createTask1, createTask2,
				},
				Count: 2,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := client.List(context.Background(), test.input)
			if err != nil {
				t.Fatalf("Failed list task err:%v", err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("%s: expected:%v, got:%v", test.name, test.want, got)
			}
		})
	}
	client.Delete(context.Background(), &pb.ByIdReq{Id: createTask1.Id})
	client.Delete(context.Background(), &pb.ByIdReq{Id: createTask2.Id})
}
