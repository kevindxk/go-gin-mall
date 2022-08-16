/**
 * @Author: dexukong
 * @Description:
 * @File:  cart
 * @Version: 1.0.0
 * @Date: 2022/08/16 9:05
 */

package model

import "gorm.io/gorm"

// 购物车模型
type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint `gorm:"not null"`
	BossID    uint
	Num       uint
	MaxNum    uint
	Check     bool
}
