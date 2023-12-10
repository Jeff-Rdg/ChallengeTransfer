package main

import (
	"ChallengeBackEndPP/configs"
	"ChallengeBackEndPP/internal/handler"
	"ChallengeBackEndPP/internal/repository"
	"ChallengeBackEndPP/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := configs.LoadDatabase()
	if err != nil {
		panic(err)
	}

	userDb := repository.NewUser(db)
	userService := user.NewService(userDb)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/user", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
	})

	http.ListenAndServe(":8000", r)
}
