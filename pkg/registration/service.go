package registration

import "github.com/marlinps/registration-payment-service/pkg/entities"

type RegistrationRepo interface {
	Create(registration entities.Registration) error
	// GetAll() ([]entities.Registration, error)
	// GetByID(registrationID string) (entities.Registration, error)
	// UpdateByID(registrationID string) (entities.Registration, error)
	// GetByEventID(eventID string) ([]entities.Registration, error)
	// GetByUserID(userID string) ([]entities.Registration, error)
	// CancelRegistration(registrationID string) error
	// GetByRegistrationID(registrationID string) (entities.Registration, error) // Added method
}

type RegistrationService struct {
	repo RegistrationRepo
}

func NewRegistrationService(repo RegistrationRepo) *RegistrationService {
	return &RegistrationService{repo: repo}
}

func (s *RegistrationService) CreateRegistration(registration entities.Registration) error {
	return s.repo.Create(registration)
}

// func (s *RegistrationService) GetAllRegistrations() ([]entities.Registration, error) {
// 	return s.repo.GetAll()
// }

// func (s *RegistrationService) GetRegistrationByID(registrationID string) (entities.Registration, error) {
// 	return s.repo.GetByID(registrationID)
// }

// func (s *RegistrationService) UpdateRegistrationByID(registrationID string) (entities.Registration, error) {
// 	return s.repo.UpdateByID(registrationID)
// }
// func (s *RegistrationService) GetRegistrationsByEventID(eventID string) ([]entities.Registration, error) {
// 	return s.repo.GetByEventID(eventID)
// }

// func (s *RegistrationService) GetRegistrationsByUserID(userID string) ([]entities.Registration, error) {
// 	return s.repo.GetByUserID(userID)
// }

// func (s *RegistrationService) CancelRegistration(registrationID string) error {
// 	return s.repo.CancelRegistration(registrationID)
// }
