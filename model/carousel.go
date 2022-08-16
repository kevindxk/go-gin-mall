/**
 * @Author: dexukong
 * @Description:
 * @File:  carousel
 * @Version: 1.0.0
 * @Date: 2022/08/16 9:02
 */

package model

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductID uint `gorm:"not null"`
}
