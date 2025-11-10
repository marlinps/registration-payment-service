package handlers

import (
	"github.com/marlinps/registration-payment-service/pkg/entities"
	"github.com/marlinps/registration-payment-service/pkg/registration"

	"github.com/gofiber/fiber/v2"
)

func CreateNewRegistration(service *registration.RegistrationService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var registration entities.Registration
		if err := c.BodyParser(&registration); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		err := service.CreateRegistration(registration)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create registration",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(registration)
	}
}

func LoadAllRegistrations(service *registration.RegistrationService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		registrations, err := service.GetAllRegistrations()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to load registrations",
			})
		}
		return c.JSON(fiber.Map{
			"data": registrations,
		})
	}
}

func LoadRegistrationByID(service *registration.RegistrationService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		registrationID := c.Params("id")
		registration, err := service.GetRegistrationByID(registrationID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to load registration",
			})
		}
		return c.JSON(registration)
	}
}
func UpdateRegistrationByID(service *registration.RegistrationService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		registrationID := c.Params("id")
		registration, err := service.UpdateRegistrationByID(registrationID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update registration",
			})
		}

		return c.JSON(registration)
	}
}

func CancelRegistration(service *registration.RegistrationService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		registrationID := c.Params("id")
		err := service.CancelRegistration(registrationID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to cancel registration",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Registration deleted successfully",
		})
	}
}
