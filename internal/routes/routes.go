package routes

import (
	"ChallengeBackEndPP/internal/handlers"
	"ChallengeBackEndPP/internal/repository"
	"ChallengeBackEndPP/transfer"
	"ChallengeBackEndPP/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
)

func LoadRoutes(db *gorm.DB) *chi.Mux {
	userDb := repository.NewUser(db)
	userService := user.NewService(userDb)
	userHandler := handlers.NewUserHandler(userService)

	transferDb := repository.NewTransfer(db)
	transferService := transfer.NewService(transferDb, userDb, db)
	transferHandler := handlers.NewTransferHandler(transferService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/user", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/{id}", userHandler.FindUserById)
		r.Put("/{id}", userHandler.AddMoney)
	})

	r.Route("/transfer", func(r chi.Router) {
		r.Post("/", transferHandler.CreateTransfer)
	})

	return r
}
