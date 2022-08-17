/**
 * @Author: dexukong
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/08/16 16:16
 */

package api

import (
	"ginmall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserServer
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
