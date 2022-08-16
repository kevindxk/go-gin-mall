/**
 * @Author: dexukong
 * @Description:
 * @File:  notice
 * @Version: 1.0.0
 * @Date: 2022/08/16 9:09
 */

package model

import "gorm.io/gorm"

// Notice 公告模型 存放公告和邮件模板
type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}
