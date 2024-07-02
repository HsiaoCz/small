package types

import "gorm.io/gorm"

// small user
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	PasswordDigest string
	Nickname       string `gorm:"unique"`
	Status         string
	Limit          int
	Avatar         string `gorm:"size:1000"`
}
