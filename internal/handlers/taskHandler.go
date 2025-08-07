package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"simpleTaskCrud/internal/data"
)

type TaskHandler struct {
	log *log.Logger
}

func NewTask(l *log.Logger) *TaskHandler {
	return &TaskHandler{l}
}

var (
	TaskRe, _     = regexp.Compile(`^/tasks/*$`)
	TaskWithId, _ = regexp.Compile(`^/tasks/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
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
	fmt.Println(r.Body)
}
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	lp := data.GetTasks()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to encode json", http.StatusInternalServerError)
	}
}
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request)    {}
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {}
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {}
