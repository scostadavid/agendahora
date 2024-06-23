package entities

import "gorm.io/gorm"

type WaitlistEntry struct {
	gorm.Model
	Email string `gorm:"uniqueIndex"`
}
