package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethod string
type VerificationStatus string

const (
	PaymentBankTransfer PaymentMethod = "bank_transfer"
	PaymentEwallet      PaymentMethod = "ewallet"
	PaymentCash         PaymentMethod = "cash"
	PaymentOther        PaymentMethod = "other"

	VerificationPending  VerificationStatus = "pending"
	VerificationApproved VerificationStatus = "approved"
	VerificationRejected VerificationStatus = "rejected"
)

type Payment struct {
	PaymentID            uuid.UUID          `gorm:"type:char(36);primaryKey" json:"payment_id"`
	RegistrationID       uuid.UUID          `gorm:"type:char(36);not null;index:idx_payments_registration" json:"registration_id"`
	Amount               float64            `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentMethod        PaymentMethod      `gorm:"type:enum('bank_transfer','ewallet','cash','other');default:'bank_transfer'" json:"payment_method"`
	PaymentDate          time.Time          `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"payment_date"`
	PaymentProofURL      string             `gorm:"type:varchar(500)" json:"payment_proof_url"`
	PaymentProofFilename string             `gorm:"type:varchar(255)" json:"payment_proof_filename"`
	BankName             string             `gorm:"type:varchar(100)" json:"bank_name"`
	AccountNumber        string             `gorm:"type:varchar(50)" json:"account_number"`
	AccountHolderName    string             `gorm:"type:varchar(255)" json:"account_holder_name"`
	VerificationStatus   VerificationStatus `gorm:"type:enum('pending','approved','rejected');default:'pending';index:idx_payments_verification_status" json:"verification_status"`
	VerifiedBy           *uuid.UUID         `gorm:"type:char(36)" json:"verified_by"`
	VerifiedAt           *time.Time         `gorm:"type:timestamp null" json:"verified_at"`
	VerificationNotes    *string            `gorm:"type:text" json:"verification_notes"`
	RejectionReason      *string            `gorm:"type:text" json:"rejection_reason"`
	CreatedAt            time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time          `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi ke Registration
	Registration Registration `gorm:"foreignKey:RegistrationID;constraint:OnDelete:CASCADE;" json:"registration"`
}

// BeforeCreate: Generate UUID otomatis
func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if p.PaymentID == uuid.Nil {
		p.PaymentID = uuid.New()
	}
	return
}
