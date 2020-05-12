package memstore

import (
	"context"
	"strconv"
	"sync"

	"github.com/romanitalian/go_server_hello/internal/models"
)

type MemStore struct {
	mu sync.Mutex
	tasks models.TasksList
	lastID int64
}

func New() *MemStore {
	return &MemStore{}
}


func (s *MemStore) GetTaskList(ctx context.Context) (models.TasksList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := models.TasksList{
		Tasks: append([]models.Task(nil), s.tasks.Tasks...),
	}

	return tasks, nil
}
func (s *MemStore) CreateTask(ctx context.Context, task models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++
	task.ID = strconv.FormatInt(s.lastID, 16)
	s.tasks.Tasks = append(s.tasks.Tasks, task)

	return task, nil

}
func (s *MemStore) DeleteTask(ctx context.Context, id string) (error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for  i, t := range s.tasks.Tasks {
		if t.ID == id {
			s.tasks.Tasks = append(s.tasks.Tasks[:i], s.tasks.Tasks[i+1:]...)
			break
		}
	}
	return nil
}