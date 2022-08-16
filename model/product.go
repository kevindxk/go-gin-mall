/**
 * @Author: dexukong
 * @Description:
 * @File:  product
 * @Version: 1.0.0
 * @Date: 2022/08/16 9:06
 */

package model

import "gorm.io/gorm"

// 商品模型
type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryID    uint   `gorm:"not null"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        int
	BossName      string
	BossAvatar    string
}
