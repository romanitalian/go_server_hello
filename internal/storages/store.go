package storages

import (
	"context"

	"github.com/romanitalian/go_server_hello/internal/models"
)

type Store interface {
	GetTaskList(ctx context.Context) (models.TasksList, error)
	CreateTask(ctx context.Context, task models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, id string) error
}