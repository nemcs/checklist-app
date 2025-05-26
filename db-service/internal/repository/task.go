package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/nemcs/checklist-app/db-service/internal/models"
)

type TaskRepository struct {
	conn *pgx.Conn
}

func NewTaskRepository(conn *pgx.Conn) *TaskRepository {
	return &TaskRepository{conn: conn}
}

func (r *TaskRepository) NewTask(ctx context.Context, task models.Task) error {
	query := `insert into checklist (id, title, description, done, created_at) values ($1, $2, $3, $4, $5)`
	_, err := r.conn.Exec(ctx, query, task.ID, task.Title, task.Description, task.Done, task.CreatedAt)
	return err
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, id string) (models.Task, error) {
	var task models.Task
	if err := r.conn.QueryRow(ctx, "select id, title, description, done, created_at from checklist where id = $1", id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Done,
		&task.CreatedAt,
	); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) GetAllTask(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task

	rows, err := r.conn.Query(ctx, "select id, title, description, done, created_at from checklist")
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Done,
			&task.CreatedAt,
		); err != nil {
			return []models.Task{}, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) UpdateDoneByID(ctx context.Context, id string) error {
	_, err := r.conn.Exec(ctx, "update checklist set done = true where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) DeleteTaskByID(ctx context.Context, id string) error {
	_, err := r.conn.Exec(ctx, "delete from checklist where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
