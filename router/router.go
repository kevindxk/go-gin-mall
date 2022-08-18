/**
 * @Author: dexukong
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/08/16 16:08
 */

package router

import (
	"ginmall/api/v1"
	"ginmall/middlerware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlerware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "succcess")
		})
		v1.POST("user/register", api.UserRegister)
		v1.GET("user/login", api.UserLogin)
	}
	return r
}
