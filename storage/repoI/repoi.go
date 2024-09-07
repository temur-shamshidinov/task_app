package repoI

import (
	"context"

	"github.com/temur-shamshidinov/task_app/models"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) ([]models.User, error)
}

type TaskRepoI interface {
	CreateTask(ctx context.Context, task models.Task) error
	GetTasks(ctx context.Context, 	id int) ([]models.Task, error)
	UpdateTask(ctx context.Context, id int, task models.Task) error
	DeleteTask(ctx context.Context, id int) error
}
