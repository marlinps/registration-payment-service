package entities

import (
	"time"
)

type Payment struct {
	PaymentID            string     `gorm:"type:char(36);primaryKey;default:(UUID())" json:"payment_id"`
	RegistrationID       string     `gorm:"type:char(36);not null;index:idx_payments_registration" json:"registration_id"`
	Amount               float64    `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentMethod        string     `gorm:"type:enum('bank_transfer','ewallet','cash','other');default:'bank_transfer'" json:"payment_method"`
	PaymentDate          time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"payment_date"`
	PaymentProofURL      string     `gorm:"type:varchar(500)" json:"payment_proof_url"`
	PaymentProofFilename string     `gorm:"type:varchar(255)" json:"payment_proof_filename"`
	BankName             string     `gorm:"type:varchar(100)" json:"bank_name"`
	AccountNumber        string     `gorm:"type:varchar(50)" json:"account_number"`
	AccountHolderName    string     `gorm:"type:varchar(255)" json:"account_holder_name"`
	VerificationStatus   string     `gorm:"type:enum('pending','approved','rejected');default:'pending';index:idx_payments_verification_status" json:"verification_status"`
	VerifiedBy           *string    `gorm:"type:char(36)" json:"verified_by"`
	VerifiedAt           *time.Time `gorm:"type:timestamp null" json:"verified_at"`
	VerificationNotes    *string    `gorm:"type:text" json:"verification_notes"`
	RejectionReason      *string    `gorm:"type:text" json:"rejection_reason"`
	CreatedAt            time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"`

	// Optional: Relasi ke model Registration jika ada
	// Registration Registration `gorm:"foreignKey:RegistrationID;constraint:OnDelete:CASCADE;" json:"registration"`
}

func (Payment) TableName() string {
	return "payments"
}
