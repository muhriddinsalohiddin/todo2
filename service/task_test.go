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
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Create(context.Background(), &tc.input)
			got.Id = ""
			got.CreatedAt = ""
			got.DeletedAt = ""
			got.UpdatedAt = ""
			if err != nil {
				t.Error("Failed create tasks", err)
			}
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: Expected:%v, got:%s", tc.name, tc.want, got)
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
	for _, tc := range tests {
		got, err := client.Get(context.Background(), &tc.input)
		if err != nil {
			t.Error("Failed get task", err)
		}
		got.Id = ""
		got.UpdatedAt = ""
		got.CreatedAt = ""
		if !reflect.DeepEqual(tc.want, *got) {
			t.Fatalf("%s: Expected:%v, \n\ngot:%s", tc.name, tc.want, got)
		}
	}
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
			name:  "update task",
			input: *task,
			want:  *task,
		},
	}
	for _, tc := range tests {
		got, err := client.Update(context.Background(), &tc.input)
		if err != nil {
			t.Error("Failed update task", err)
		}
		if reflect.DeepEqual(tc.want.UpdatedAt, got.UpdatedAt) {
			t.Fatalf("%s: Expected:%v, \n\ngot:%s", tc.name, tc.want.UpdatedAt, got.UpdatedAt)
		}
	}
}
