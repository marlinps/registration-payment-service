package registration

import (
	"github.com/marlinps/registration-payment-service/pkg/entities"

	"gorm.io/gorm"
)

type RegistrationRepository struct {
	db *gorm.DB
}

func NewRegistrationRepository(db *gorm.DB) *RegistrationRepository {
	return &RegistrationRepository{
		db: db,
	}
}
func (r *RegistrationRepository) Create(registration entities.Registration) error {
	return r.db.Create(&registration).Error
}

func (r *RegistrationRepository) GetAll() ([]entities.Registration, error) {
	var registrations []entities.Registration
	err := r.db.Preload("Payments").Find(&registrations).Error
	return registrations, err
}

// func (r *RegistrationRepository) GetByID(registrationID string) (entities.Registration, error) {
// 	var registration entities.Registration
// 	err := r.db.Preload("Payments").Where("registration_id = ?", registrationID).First(&registration).Error
// 	return registration, err
// }

// func (r *RegistrationRepository) GetByRegistrationID(registrationID string) (entities.Registration, error) {
// 	var registration entities.Registration
// 	err := r.db.Preload("Payments").Where("registration_id = ?", registrationID).First(&registration).Error
// 	return registration, err
// }

// func (r *RegistrationRepository) UpdateByID(registrationID string) (entities.Registration, error) {
// 	var registration entities.Registration
// 	err := r.db.Where("registration_id = ?", registrationID).First(&registration).Error
// 	if err != nil {
// 		return registration, err
// 	}
// 	return registration, r.db.Save(&registration).Error
// }

// func (r *RegistrationRepository) GetByEventID(eventID string) ([]entities.Registration, error) {
// 	var registrations []entities.Registration
// 	err := r.db.Preload("Payments").Where("event_id = ?", eventID).Find(&registrations).Error
// 	return registrations, err
// }

// func (r *RegistrationRepository) GetByUserID(userID string) ([]entities.Registration, error) {
// 	var registrations []entities.Registration
// 	err := r.db.Preload("Payments").Where("user_id = ?", userID).Find(&registrations).Error
// 	return registrations, err
// }

// func (r *RegistrationRepository) CancelRegistration(registrationID string) error {

// 	return r.db.Where("registration_id = ?", registrationID).Delete(&entities.Registration{}).Error
// }
