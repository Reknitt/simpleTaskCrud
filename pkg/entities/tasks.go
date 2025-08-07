package entities

import (
	"encoding/json"
	"io"
)

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
}

type Tasks []*Task

func (t *Tasks) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}
