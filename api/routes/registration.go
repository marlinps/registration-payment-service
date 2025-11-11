package routes

import (
	"github.com/marlinps/registration-payment-service/api/handlers"
	"github.com/marlinps/registration-payment-service/pkg/registration"

	"github.com/gofiber/fiber/v2"
)

func RegistrationRoutes(api fiber.Router, registrationService *registration.RegistrationService) {
	api.Post("/registrations", handlers.CreateNewRegistration(registrationService))
	// api.Get("/registrations", handlers.LoadAllRegistrations(registrationService))
	// api.Get("/registrations/:id", handlers.LoadRegistrationByID(registrationService))
	// api.Put("/registrations/:id", handlers.UpdateRegistrationByID(registrationService))
	// api.Delete("/registrations/:id", handlers.CancelRegistration(registrationService))
}
