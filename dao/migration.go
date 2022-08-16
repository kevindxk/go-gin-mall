/**
 * @Author: dexukong
 * @Description:
 * @File:  migration
 * @Version: 1.0.0
 * @Date: 2022/08/15 15:49
 */

package dao

import (
	"fmt"
	"ginmall/model"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Carousel{},
		&model.Cart{},
		&model.Favorite{},
		&model.Notice{},
		&model.Product{},
		&model.ProductImg{},
		&model.Admin{},
		&model.Adress{},
		&model.Order{})
	if err != nil {
		fmt.Println("err", err)
	}
	return
}
