package tasks

import (
	"encoding/json"
	"errors"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Done     bool   `json:"done"`
	Priority string `json:"priority"`
}

var tasksFile = "pkg/tasks/tasks.json"

func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(tasksFile); os.IsNotExist(err) {
		return []Task{}, nil
	}
	data, err := os.ReadFile(tasksFile)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFile, data, 0644)
}

func AddTask(title string, priority string) error {
	tasks, _ := LoadTasks()
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	tasks = append(tasks, Task{ID: id, Title: title, Done: false, Priority: priority})
	return SaveTasks(tasks)
}

func GetTask(id int) bool {
	tasks, _ := LoadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			return true
		}
	}
	return false
}

func CompleteTask(id int) error {
	tasks, _ := LoadTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			// reassign the new modified task here
			return SaveTasks(tasks)
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id int) ([]Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return []Task{}, errors.New("Error while loading the task")
	}
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			SaveTasks(tasks)
			return tasks, nil
		}
	}
	return []Task{}, errors.New("Particular task not found")
}
