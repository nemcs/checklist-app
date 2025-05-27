package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nemcs/checklist-app/db-service/internal/models"
)

type ChecklistRepo struct {
	pool *pgxpool.Pool
}

func New(dbpool *pgxpool.Pool) *ChecklistRepo {
	return &ChecklistRepo{pool: dbpool}
}

func (r *ChecklistRepo) NewTask(ctx context.Context, task models.Task) error {
	query := "insert into checklist (id, title, description) values ($1, $2, $3)"
	commandTag, err := r.pool.Exec(ctx, query, task.ID, task.Title, task.Description)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("Задача не создана")
	}
	return nil
}

func (r *ChecklistRepo) GetTaskByID(ctx context.Context, id string) (models.Task, error) {
	var task models.Task
	query := "select id, title, description, done, created_at from checklist where id = $1"
	if err := r.pool.QueryRow(ctx, query, id).Scan(
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

func (r *ChecklistRepo) GetAllTask(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	query := "select id, title, description, done, created_at from checklist"
	rows, err := r.pool.Query(ctx, query)
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

func (r *ChecklistRepo) UpdateDoneByID(ctx context.Context, id string) error {
	query := "update checklist set done = true where id = $1"
	commandTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("Задача %s не найдена", id)
	}
	return nil
}

func (r *ChecklistRepo) DeleteTaskByID(ctx context.Context, id string) error {
	query := "delete from checklist where id = $1"
	commandTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("Задача с %s не найдена", id)
	}
	return nil
}
