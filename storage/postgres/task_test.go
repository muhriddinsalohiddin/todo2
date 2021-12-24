package postgres

import (
	"reflect"
	"testing"

	pb "github.com/muhriddinsalohiddin/todo2/genproto"
)

func TestTaskRepo_Create(t *testing.T) {
	tests := []struct {
		name  string
		input pb.Task
		want  pb.Task
	}{
		{
			name: "Create",
			input: pb.Task{
				Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e5",
				Assignee: "Assignee",
				Title:    "Title",
				Summary:  "Summary",
				Deadline: "2022-04-04",
				Status:   "Aktive",
			},
			want: pb.Task{
				Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e5",
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
			got, err := pgRepo.Create(test.input)
			if err != nil {
				t.Fatalf("Failed create user err:%v", err)
			}
			got.UpdatedAt = ""
			got.CreatedAt = ""
			if reflect.DeepEqual(test.input, got) {
				t.Fatalf("%s: expected:%v, got:%v", test.name, test.want, got)
			}
		})
	}
}
func TestTaskRepo_Get(t *testing.T) {
	createtask, err := pgRepo.Create(pb.Task{
		Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e5",
		Assignee: "Assignee",
		Title:    "Title",
		Summary:  "Summary",
		Deadline: "2022-04-04",
		Status:   "Aktive",
	})
	if err != nil {
		t.Fatalf("Failed create user err:%v", err)
	}
	tests := []struct {
		name  string
		input string
		want  pb.Task
	}{
		{
			name:  "GET USER",
			input: createtask.Id,
			want:  createtask,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := pgRepo.Get(test.input)
			if err != nil {
				t.Fatalf("Failed get user err:%v", err)
			}
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("%s: expected:%v,\n\n\n\n got:%v", test.name, test.want, got)
			}
		})
	}
}
