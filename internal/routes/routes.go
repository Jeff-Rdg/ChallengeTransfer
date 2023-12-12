package routes

import (
	"ChallengeBackEndPP/internal/handlers"
	"ChallengeBackEndPP/internal/repository"
	"ChallengeBackEndPP/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
)

func LoadRoutes(db *gorm.DB) *chi.Mux {
	userDb := repository.NewUser(db)
	userService := user.NewService(userDb)
	userHandler := handlers.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/user", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/{id}", userHandler.FindUserById)
	})

	return r
}
