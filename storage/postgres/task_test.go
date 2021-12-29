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
				t.Fatalf("Failed create task err:%v", err)
			}
			got.UpdatedAt = ""
			got.CreatedAt = ""
			if reflect.DeepEqual(test.input, got) {
				t.Fatalf("%s: expected:%v, got:%v", test.name, test.want, got)
			}
		})
		pgRepo.Delete(test.input.Id)
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
		t.Fatalf("Failed create task err:%v", err)
	}
	tests := []struct {
		name  string
		input string
		want  pb.Task
	}{
		{
			name:  "GET TASK",
			input: createtask.Id,
			want:  createtask,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := pgRepo.Get(test.input)
			if err != nil {
				t.Fatalf("Failed get task err:%v", err)
			}
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("%s: expected:%v, got:%v", test.name, test.want, got)
			}
		})
		pgRepo.Delete(test.want.Id)
	}
}

func TestTaskRepo_Update(t *testing.T) {
	createTest, _ := pgRepo.Create(pb.Task{
		Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e5",
		Assignee: "Assignee",
		Title:    "Title",
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
			name:  "Update test",
			input: createTest,
			want:  createTest,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := pgRepo.Update(test.input)
			if err != nil {
				t.Fatalf("Failed update task err:%v", err)
			}
			if reflect.DeepEqual(test.want.UpdatedAt, got.UpdatedAt) {
				t.Fatalf("%s: expected:%v, got:%v", test.name, test.want.UpdatedAt, got.UpdatedAt)
			}
		})
		pgRepo.Delete(test.want.Id)
	}
}

func TestTaskRepo_List(t *testing.T) {
	createTask1, _ := pgRepo.Create(pb.Task{
		Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e5",
		Assignee: "Assignee",
		Title:    "Title",
		Summary:  "Summary",
		Deadline: "2022-04-03",
		Status:   "Aktive",
	})
	createTask2, _ := pgRepo.Create(pb.Task{
		Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e2",
		Assignee: "Assignee",
		Title:    "Title",
		Summary:  "Summary",
		Deadline: "2022-01-04",
		Status:   "Aktive",
	})

	tests := []struct {
		name  string
		input pb.ListReq
		want  pb.ListResp
	}{
		{
			name: "List",
			input: pb.ListReq{
				Page:  1,
				Limit: 10,
			},
			want: pb.ListResp{
				Tasks: []*pb.Task{
					&createTask1, &createTask2,
				},
				Count: 2,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := pgRepo.List(test.input)
			if err != nil {
				t.Fatalf("Failed list task err:%v", err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("%s: expected:%v, got:%v", test.name, test.want, got)
			}
		})
	}
	pgRepo.Delete(createTask1.Id)
	pgRepo.Delete(createTask2.Id)
}

func TestTaskRepo_Delete(t *testing.T) {
	createTask, _ := pgRepo.Create(pb.Task{
		Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e5",
		Assignee: "Assignee",
		Title:    "Title",
		Summary:  "Summary",
		Deadline: "2022-04-04",
		Status:   "Aktive",
	})
	tests := []struct {
		name  string
		input string
		want  error
	}{
		{
			name:  "Delete task",
			input: createTask.Id,
			want:  nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := pgRepo.Delete(test.input)
			if err != nil {
				t.Fatalf("Failed delete task err:%v", err)
			}
		})
	}
}

func TestTaskRepo_ListOverdue(t *testing.T) {
	createTask1, _ := pgRepo.Create(pb.Task{
		Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e5",
		Assignee: "Assignee",
		Title:    "Title",
		Summary:  "Summary",
		Deadline: "2022-04-03",
		Status:   "Aktive",
	})
	createTask2, _ := pgRepo.Create(pb.Task{
		Id:       "39bf6bdf-2e62-4687-bbc6-2bc1abd626e2",
		Assignee: "Assignee",
		Title:    "Title",
		Summary:  "Summary",
		Deadline: "2022-01-04",
		Status:   "Aktive",
	})

	tests := []struct {
		name  string
		input pb.ListOverReq
		want  pb.ListOverResp
	}{
		{
			name: "ListOverdue",
			input: pb.ListOverReq{
				Time:  "2020-04-04",
				Page:  1,
				Limit: 10,
			},
			want: pb.ListOverResp{
				Tasks: []*pb.Task{
					&createTask1, &createTask2,
				},
				Count: 2,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := pgRepo.ListOverdue(test.input)
			if err != nil {
				t.Fatalf("Failed list Overdue task err:%v", err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("%s: expected:%v, got:%v", test.name, test.want, got)
			}
		})
	}
	pgRepo.Delete(createTask1.Id)
	pgRepo.Delete(createTask2.Id)
}
