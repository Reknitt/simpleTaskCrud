package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "Hello %s", name)
}
