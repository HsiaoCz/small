package types

import "gorm.io/gorm"

// ProductParamImg 商品图片模型
type ProductInfoImg struct {
	gorm.Model
	ProductID uint
	ImgPath   string
}
