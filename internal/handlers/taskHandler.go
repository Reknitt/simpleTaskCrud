package handlers

import (
	"log"
	"net/http"
	"regexp"
	"simpleTaskCrud/internal/data"
	"simpleTaskCrud/pkg/entities"
	"strconv"
)

type TaskHandler struct {
	log *log.Logger
}

func NewTask(l *log.Logger) *TaskHandler {
	return &TaskHandler{l}
}

var (
	TaskRe, _     = regexp.Compile(`^/tasks/*$`)
	TaskWithId, _ = regexp.Compile(`/\d+`)
)

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle http methods
	switch {
	case r.Method == http.MethodPost && TaskRe.MatchString(r.URL.Path):
		h.CreateTask(w, r)
		return
	case r.Method == http.MethodGet && TaskRe.MatchString(r.URL.Path):
		h.ListTasks(w, r)
		return
	case r.Method == http.MethodGet && TaskWithId.MatchString(r.URL.Path):
		h.GetTask(w, r)
		return
	case r.Method == http.MethodPut && TaskWithId.MatchString(r.URL.Path):
		h.UpdateTask(w, r)
		return
	case r.Method == http.MethodDelete && TaskWithId.MatchString(r.URL.Path):
		h.DeleteTask(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	task := &entities.Task{}
	err := task.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to decode json", http.StatusBadRequest)
		return
	}
	data.CreateTask(task)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	lp := data.GetTasks()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to encode json", http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	f := TaskWithId.FindAllString(r.URL.Path, -1)
	if len(f) != 1 {
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}
	idString := f[0][1:]

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Bad Id", http.StatusBadRequest)
		return
	}

	ent, err := data.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ent.ToJSON(w)
	if err != nil {
		http.Error(w, "unable to encode json", http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {}
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {}
