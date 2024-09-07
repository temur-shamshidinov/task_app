package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/task_app/models"
	"github.com/temur-shamshidinov/task_app/storage/repoI"
)

type taskRepo struct {
	con *pgx.Conn
}

func NewTaskRepo(con *pgx.Conn) repoI.TaskRepoI {
	return &taskRepo{
		con: con,
	}
}

func (t *taskRepo) CreateTask(ctx context.Context, task models.Task) error {
	query := `
		INSERT INTO tasks (
			user_id,
			title
		) VALUES (
			$1, $2
		)
	`

	_, err := t.con.Exec(
		ctx,
		query,
		task.UserID,
		task.Title,
	)

	if err != nil {
		log.Println("Error inserting task:", err)
		return err
	}

	return nil
}

func (t *taskRepo) GetTasks(ctx context.Context, userID int) ([]models.Task, error) {

	var tasks []models.Task

	query := `
		SELECT 
			id, 
			user_id, 
			title, 
			created_at
		FROM 
			tasks
		WHERE
			user_id = $1`

	rows, err := t.con.Query(ctx, query, userID)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var task models.Task

		if err := rows.Scan(
			&task.ID, 
			&task.UserID, 
			&task.Title, 
			&task.CreatedAt,
		); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error retrieving rows:", err)
		return nil, err
	}
	return tasks, nil
}


func (t *taskRepo) UpdateTask(ctx context.Context, id int, task models.Task) error {
    query := `
        UPDATE tasks
        SET user_id = $1, title = $2
        WHERE id = $3`
    _, err := t.con.Exec(ctx, query, task.UserID, task.Title, id)
    if err != nil {
        log.Println("Error executing query:", err)
        return err
    }
    return nil
}

func (t *taskRepo) DeleteTask(ctx context.Context, id int) error {
    query := `
        DELETE FROM tasks 
        WHERE id = $1`
    _, err := t.con.Exec(ctx, query, id)
    if err != nil {
        log.Println("Error executing query:", err)
        return err
    }
    return nil
}

