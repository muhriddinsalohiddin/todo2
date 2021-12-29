package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pb "github.com/muhriddinsalohiddin/todo2/genproto"
)

type taskRepo struct {
	db *sqlx.DB
}

// NewTaskRezzpo
func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (t *taskRepo) Create(task pb.Task) (pb.Task, error) {
	err := t.db.QueryRow(
		`INSERT INTO tasks (id, assignee, title, summary, deadline, status, created_at,updated_at)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`,
		task.Id,
		task.Assignee,
		task.Title,
		task.Summary,
		task.Deadline,
		task.Status,
		time.Now().UTC(),
		time.Now().UTC(),
	).Scan(&task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	task, err = t.Get(task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (t *taskRepo) Get(id string) (pb.Task, error) {
	var task pb.Task
	err := t.db.QueryRow(`
		SELECT id,
			assignee,
			title,
			summary,
			deadline,
			status,
			created_at,
			updated_at
		FROM tasks
		WHERE id = $1
		AND deleted_at IS NULL`,
		id).Scan(
		&task.Id,
		&task.Assignee,
		&task.Title,
		&task.Summary,
		&task.Deadline,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (t *taskRepo) List(l pb.ListReq) (pb.ListResp, error) {
	offset := (l.Page - 1) * l.Limit

	rows, err := t.db.Queryx(`
		SELECT id,
			assignee,
			title,
			SUMMARY,
			deadline,
			status,
			created_at,
			updated_at
		FROM tasks
		WHERE deleted_at IS NULL
		LIMIT $1
		OFFSET $2`,
		l.Limit, offset)
	if err != nil {
		return pb.ListResp{}, err
	}
	if err = rows.Err(); err != nil {
		return pb.ListResp{}, err
	}
	defer rows.Close()

	var list pb.ListResp
	for rows.Next() {
		var task pb.Task
		err = rows.Scan(
			&task.Id,
			&task.Assignee,
			&task.Title,
			&task.Summary,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return pb.ListResp{}, err
		}
		list.Tasks = append(list.Tasks, &task)
	}
	err = t.db.QueryRow(`select count(*) from tasks`).Scan(&list.Count)
	if err != nil {
		return pb.ListResp{}, err
	}
	return list, nil
}

func (t *taskRepo) Update(task pb.Task) (pb.Task, error) {
	result, err := t.db.Exec(`
		UPDATE tasks
		SET assignee=$1,
			title=$2,
			summary=$3,
			deadline=$4,
			status=$5,
			updated_at=$6
		WHERE id=$7`,
		task.Assignee,
		task.Title,
		task.Summary,
		task.Deadline,
		task.Status,
		time.Now().UTC(),
		task.Id)
	if err != nil {
		return pb.Task{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Task{}, sql.ErrNoRows
	}

	task, err = t.Get(task.Id)
	if err != nil {
		return pb.Task{}, err
	}
	return task, nil
}

func (t *taskRepo) Delete(id string) error {
	result, err := t.db.Exec(`update tasks set deleted_at=$1 where id = $2`, time.Now().UTC(), id)
	if err != nil {
		return err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (t *taskRepo) ListOverdue(l pb.ListOverReq) (pb.ListOverResp, error) {
	offset := (l.Page - 1) * l.Limit
	duration, err := time.Parse("2006-01-02", l.Time)
	if err != nil {
		return pb.ListOverResp{}, err
	}

	rows, err := t.db.Query(`
		SELECT 
			id,
			assignee,
			title,
			SUMMARY,
			deadline,
			status,
			created_at,
			updated_at
		FROM tasks
		WHERE deadline > $1
		AND deleted_at IS NULL
		LIMIT $2
		OFFSET $3`,
		duration,
		l.Limit,
		offset,
	)
	if err != nil {
		return pb.ListOverResp{}, nil
	}
	var list pb.ListOverResp
	for rows.Next() {
		var task pb.Task
		err = rows.Scan(
			&task.Id,
			&task.Assignee,
			&task.Title,
			&task.Summary,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return pb.ListOverResp{}, nil
		}
		list.Tasks = append(list.Tasks, &task)
	}
	err = t.db.QueryRow(`select count(*) from tasks where deadline > $1`,
		duration).Scan(&list.Count)
	if err != nil {
		return pb.ListOverResp{}, nil
	}
	return list, nil
}
