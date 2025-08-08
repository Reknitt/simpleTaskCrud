package data

import (
	"fmt"
	"simpleTaskCrud/pkg/entities"
	"slices"
	"time"
)

func GetTasks() entities.Tasks {
	return tasks
}

func CreateTask(task *entities.Task) {
	task.Id = getId()
	tasks = append(tasks, task)
}

func GetTask(id int) (*entities.Task, error) {
	idx := slices.IndexFunc(tasks, func(t *entities.Task) bool {
		return t.Id == id
	})
	if idx == -1 {
		return nil, fmt.Errorf("task with id %d not found", id)
	}
	return tasks[idx], nil
}

func getId() int {
	if len(tasks) == 0 {
		return 1
	}
	task := tasks[len(tasks)-1]
	return task.Id + 1
}

var tasks = []*entities.Task{
	&entities.Task{
		Id:          1,
		Name:        "Убраться",
		Description: "Необходимо убраться дома",
		Completed:   false,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	&entities.Task{
		Id:          2,
		Name:        "Постирать вещи",
		Description: "Необходимо постирать вещи",
		Completed:   false,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
