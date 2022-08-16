/**
 * @Author: dexukong
 * @Description:
 * @File:  product_img
 * @Version: 1.0.0
 * @Date: 2022/08/16 9:07
 */

package model

import "gorm.io/gorm"

type ProductImg struct {
	gorm.Model
	ProductID uint `gorm:"not null"`
	ImgPath   string
}
