package handler

import (
	"fiber-clean-arch-demo/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(u domain.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.usecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var request domain.CreateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if request.Name == "" || request.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name and email are required"})
	}

	user, err := h.usecase.Create(&request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
