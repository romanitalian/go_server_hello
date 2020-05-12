package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/romanitalian/go_server_hello/internal/models"
	"github.com/romanitalian/go_server_hello/internal/router"
	"github.com/romanitalian/go_server_hello/internal/storages/memstore"
	"github.com/stretchr/testify/require"
)

func TestNotFoud(t *testing.T) {
	rt := router.New(nil)
	srv := httptest.NewServer(rt.RootHandler())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/unknown")
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestEmptyTaskList(t *testing.T) {
	store := memstore.New()
	rt := router.New(store)
	srv := httptest.NewServer(rt.RootHandler())
	defer srv.Close()

	checkTasksList(t, srv.URL)
}

func TestCreateTask(t *testing.T) {
	store := memstore.New()
	rt := router.New(store)
	srv := httptest.NewServer(rt.RootHandler())
	defer srv.Close()

	respTask := createTask(t, srv.URL, "test task")
	checkTasksList(t, srv.URL, respTask)
}

func TestDeleteTask(t *testing.T) {
	store := memstore.New()
	rt := router.New(store)
	srv := httptest.NewServer(rt.RootHandler())
	defer srv.Close()

	respTask := createTask(t, srv.URL, "test task")
	deleteTask(t, srv.URL, respTask.ID)
	checkTasksList(t, srv.URL)
}

// helpers

func createTask(t *testing.T, url string, text string) models.Task {
	task := models.Task{
		Text: text,
	}
	data, err := json.Marshal(&task)
	require.NoError(t, err)

	resp, err := http.Post(url+"/task", "application/json", bytes.NewReader(data))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var respTask models.Task
	jsDecoder := json.NewDecoder(resp.Body)
	err = jsDecoder.Decode(&respTask)
	require.NoError(t, err)

	return respTask
}

func checkTasksList(t *testing.T, url string, tasks ...models.Task) {
	resp, err := http.Get(url + "/tasks")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var taskList models.TasksList
	jsDecoder := json.NewDecoder(resp.Body)
	err = jsDecoder.Decode(&taskList)
	require.NoError(t, err)

	require.Equal(t, tasks, taskList.Tasks)
}

func deleteTask(t *testing.T, url string, id string) {
	req, err := http.NewRequest(http.MethodDelete, url + "/task/" + id, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
