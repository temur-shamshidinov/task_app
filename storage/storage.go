package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/task_app/storage/postgres"
	"github.com/temur-shamshidinov/task_app/storage/repoI"
)

type StorageI interface {
	GetTaskRepo() repoI.TaskRepoI
	GetUserRepo() repoI.UserRepoI
}

type storage struct {
	taskRepo repoI.TaskRepoI
	userRepo repoI.UserRepoI
}

func (s *storage) GetTaskRepo() repoI.TaskRepoI {

	return s.taskRepo
}

func (s *storage) GetUserRepo() repoI.UserRepoI {

	return s.userRepo
}

func NewStorage(con *pgx.Conn) StorageI {
	return &storage{
		taskRepo: postgres.NewTaskRepo(con),
		userRepo: postgres.NewUserRepo(con),
	}
}
