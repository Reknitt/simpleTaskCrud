package data

import (
	"simpleTaskCrud/pkg/entities"
	"time"
)

func GetTasks() entities.Tasks {
	return tasks
}

var tasks = []*entities.Task{
	&entities.Task{
		Name:        "Убраться",
		Description: "Необходимо убраться дома",
		Completed:   false,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	&entities.Task{
		Name:        "Постирать вещи",
		Description: "Необходимо постирать вещи",
		Completed:   false,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
