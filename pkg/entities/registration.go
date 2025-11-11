package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ENUMS
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
	RegistrationID           uuid.UUID          `gorm:"type:char(36);primaryKey" json:"registration_id"`
	EventID                  uuid.UUID          `gorm:"type:char(36);not null;index:idx_event_user,unique;index:idx_registrations_event;index:idx_event_status,priority:1" json:"event_id"`
	UserID                   uuid.UUID          `gorm:"type:char(36);not null;index:idx_event_user,unique;index:idx_registrations_user" json:"user_id"`
	FullName                 string             `gorm:"size:255;not null" json:"full_name"`
	Gender                   Gender             `gorm:"type:enum('male','female');not null" json:"gender"`
	Phone                    string             `gorm:"size:20;not null" json:"phone"`
	Email                    string             `gorm:"size:255;not null" json:"email"`
	Address                  string             `gorm:"type:text" json:"address"`
	EmergencyContactName     string             `gorm:"size:255" json:"emergency_contact_name"`
	EmergencyContactPhone    string             `gorm:"size:20" json:"emergency_contact_phone"`
	EmergencyContactRelation string             `gorm:"size:100" json:"emergency_contact_relation"`
	SpecialNeeds             string             `gorm:"type:text" json:"special_needs"`
	RegistrationDate         time.Time          `gorm:"autoCreateTime" json:"registration_date"`
	Status                   RegistrationStatus `gorm:"type:enum('pending','paid','confirmed','cancelled','rejected');default:'pending';not null;index:idx_registrations_status;index:idx_event_status,priority:2" json:"status" default:"pending"`
	CancelledAt              *time.Time         `json:"cancelled_at"`
	CancellationReason       string             `gorm:"type:text" json:"cancellation_reason"`
	Notes                    string             `gorm:"type:text" json:"notes"`
	CreatedAt                time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt                time.Time          `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi ke Payment (one-to-many)
	// Payments []Payment `gorm:"foreignKey:RegistrationID;constraint:OnDelete:CASCADE;" json:"payments"`
}

// BeforeCreate: Generate UUID & validasi wajib
func (r *Registration) BeforeCreate(tx *gorm.DB) (err error) {
	if r.RegistrationID == uuid.Nil {
		r.RegistrationID = uuid.New()
	}

	if r.EventID == uuid.Nil {
		r.EventID = uuid.New()
	}

	if r.UserID == uuid.Nil {
		r.UserID = uuid.New()
	}

	if r.FullName == "" || r.Email == "" || r.Phone == "" {
		return fmt.Errorf("required fields: full_name, email, and phone cannot be empty")
	}

	return
}
