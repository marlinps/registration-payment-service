package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marlinps/registration-payment-service/pkg/entities"
)

type RegistrationPresenter struct {
	RegistrationID string `json:"registration_id"`
	EventID        string `json:"event_id"`
	UserID         string `json:"user_id"`
	FullName       string `json:"full_name"`
	Gender         Gender `type:enum('male','female');"json:"gender"`
	Status         string `json:"status"`
	TimeStamp      string `json:"time_stamp"`
}

func RegistrationSuccessResponse(data *entities.Registration) *fiber.Map {
	registration := RegistrationPresenter{
		RegistrationID: data.RegistrationID.String(),
		EventID:        data.EventID.String(),
		UserID:         data.UserID.String(),
		FullName:       data.FullName,
		Gender:         data.Gender,
		Status:         data.Status,
		TimeStamp:      data.TimeStamp,
	}
	return &fiber.Map{
		"status": "success",
		"data":   registration,
	}
}
