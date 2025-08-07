package main

import (
	"log"
	"net/http"
	"os"
	"simpleTaskCrud/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	l := log.New(os.Stdout, "api", log.LstdFlags)
	th := handlers.NewTask(l)
	hh := handlers.NewHello(l)
	mux.Handle("/", hh)
	mux.Handle("/tasks", th)
	http.ListenAndServe(":80", mux)
}
