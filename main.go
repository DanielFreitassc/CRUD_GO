package main

import (
	"crudgo"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		Panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}" handlers.Update)
	r.Delete("/{id}" handlers.Delete)
	r.GetAll("/" handlers.List)
	r.Get("/{id}" handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()),r)
}