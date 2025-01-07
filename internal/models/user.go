package models

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	VATNumber   string    `gorm:"type:varchar(100);uniqueIndex:idx_vat_company"`
	CompanyName *string   `gorm:"type:varchar(255);uniqueIndex:idx_vat_company"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key"`
	FirstName      string    `gorm:"type:varchar(100);not null"`
	LastName       string    `gorm:"type:varchar(100);not null"`
	Email          string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password       string    `gorm:"type:varchar(255);not null"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	Organization   Organization
	CreatedAt      time.Time
	UpdatedAt      time.Time
}