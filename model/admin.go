/**
 * @Author: dexukong
 * @Description:
 * @File:  admin
 * @Version: 1.0.0
 * @Date: 2022/08/16 9:00
 */

package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string `gorm:"size:1000"`
}
