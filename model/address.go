/**
 * @Author: dexukong
 * @Description:
 * @File:  address
 * @Version: 1.0.0
 * @Date: 2022/08/15 15:55
 */

package model

import "gorm.io/gorm"

type Adress struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Name    string `gorm:"type:varchar(20) not null"`
	Phone   string `gorm:"type:varchar(11) not null"`
	Address string `gorm:"type:varchar(50) not null"`
}
