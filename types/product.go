package types

import "gorm.io/gorm"

// product
type Product struct {
	gorm.Model
	Name          string
	CategoryID    int
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
}
