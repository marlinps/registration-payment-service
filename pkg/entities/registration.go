package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegistrationStatus string
type Gender string

const (
	StatusPending   RegistrationStatus = "pending"
	StatusPaid      RegistrationStatus = "paid"
	StatusConfirmed RegistrationStatus = "confirmed"
	StatusCancelled RegistrationStatus = "cancelled"
	StatusRejected  RegistrationStatus = "rejected"

	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

type Registration struct {
	RegistrationID           string             `gorm:"type:char(36);primaryKey" json:"registration_id"`
	EventID                  string             `gorm:"type:char(36);not null;index:idx_event_user,unique;index:idx_registrations_event;index:idx_event_status,priority:1" json:"event_id"`
	UserID                   *string            `gorm:"type:char(36);index:idx_event_user,unique;index:idx_registrations_user" json:"user_id"`
	FullName                 string             `gorm:"size:255;not null" json:"full_name"`
	Gender                   Gender             `gorm:"type:enum('male','female');not null" json:"gender"`
	Phone                    string             `gorm:"size:20;not null" json:"phone"`
	Email                    string             `gorm:"size:255;not null" json:"email"`
	Address                  string             `json:"address"`
	EmergencyContactName     string             `gorm:"size:255" json:"emergency_contact_name"`
	EmergencyContactPhone    string             `gorm:"size:20" json:"emergency_contact_phone"`
	EmergencyContactRelation string             `gorm:"size:100" json:"emergency_contact_relation"`
	SpecialNeeds             string             `json:"special_needs"`
	RegistrationDate         time.Time          `gorm:"autoCreateTime" json:"registration_date"`
	Status                   RegistrationStatus `gorm:"type:enum('pending','paid','confirmed','cancelled','rejected');default:'pending';not null;index:idx_registrations_status;index:idx_event_status,priority:2" json:"status"`
	CancelledAt              *time.Time         `json:"cancelled_at"`
	CancellationReason       string             `json:"cancellation_reason"`
	Notes                    string             `json:"notes"`
	CreatedAt                time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt                time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook untuk generate UUID & validasi
func (r *Registration) BeforeCreate(tx *gorm.DB) (err error) {
	if r.RegistrationID == "" {
		r.RegistrationID = uuid.New().String()
	}

	if r.FullName == "" || r.Email == "" || r.Phone == "" {
		return fmt.Errorf("required fields: full_name, email, and phone cannot be empty")
	}

	return
}
