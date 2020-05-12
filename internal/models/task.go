package models


type Task struct {
	ID string `json:"id"`
	Text string `json:"text"`
}

type TasksList struct {
	Tasks []Task `json:"tasks"`
}
