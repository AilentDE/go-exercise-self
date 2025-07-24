package routes

import (
	"fiber-clean-arch-demo/internal/handler"
	"fiber-clean-arch-demo/internal/repository"
	"fiber-clean-arch-demo/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	userGroup := app.Group("/users")
	userGroup.Get("/", userHandler.GetAllUsers)
	userGroup.Post("/", userHandler.CreateUser)
}