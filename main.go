package main

import (
	"baim/routes"

	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	tasks := http.NewServeMux()
	tasks.Handle("GET /{$}", routes.Middleware(http.HandlerFunc(routes.IndexHandler)))
	tasks.Handle("GET /session", routes.Middleware(http.HandlerFunc(routes.GetSessionHandler)))
	tasks.Handle("/", http.FileServer(http.Dir("static")))

	app := http.NewServeMux()
	app.Handle("/task1", routes.Middleware(http.HandlerFunc(routes.Task1Handler)))
	app.Handle("/task2", routes.Middleware(http.HandlerFunc(routes.Task2Handler)))
	app.Handle("/task3", routes.Middleware(http.HandlerFunc(routes.Task3Handler)))

	go (func() {
		wg.Add(1)
		defer wg.Done()

		http.ListenAndServe(":8080", tasks)
	})()
	go (func() {
		wg.Add(1)
		defer wg.Done()

		http.ListenAndServe(":8081", app)
	})()

	log.Println("Server started on :8080 for tasks and :8081 for app")
	wg.Wait()
}
