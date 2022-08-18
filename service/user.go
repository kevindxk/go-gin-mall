/**
 * @Author: dexukong
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/08/16 16:17
 */

package service

import (
	"context"
	"ginmall/dao"
	"ginmall/model"
	"ginmall/pkg/c"
	"ginmall/serializer"
	"ginmall/utils"
	logging "github.com/sirupsen/logrus"
)

type UserServer struct {
	NickName string `json:"nick_name" from :"nick_name"`
	UserName string `json:"user_name" from :"user_name"`
	Password string `json:"password" from :"password"`
	Key      string `json:"nick_name" from :"key"`
}

//用户注册
func (server *UserServer) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := c.Success
	if server.Key == "" || len(server.Key) != 16 {
		code = c.Error
		return serializer.Response{
			Status: code,
			Msg:    c.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}
	//---->密文加密存储
	utils.Encrypt.SetKey(server.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(server.UserName)
	if err != nil {
		code = c.Error
		return serializer.Response{
			Status: code,
			Msg:    c.GetMsg(code),
		}
	}
	if exist {
		code = c.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    c.GetMsg(code),
		}
	}
	user = model.User{
		UserName: server.UserName,
		NickName: server.NickName,
		Status:   model.ACtive,
		Avatar:   "",
		Money:    utils.Encrypt.AesEncoding("10000"),
	}
	//密码加密
	if err = user.SetPassword(server.Password); err != nil {
		code = c.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    c.GetMsg(code),
		}
	}
	//创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = c.Error
	}
	return serializer.Response{
		Status: code,
		Msg:    c.GetMsg(code),
	}
}

//用户登录
func (server *UserServer) Login(ctx context.Context) serializer.Response {
	code := c.Success
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(server.UserName)
	if !exist {
		logging.Info(err)
		code = c.ErrorUserNotFond
		return serializer.Response{
			Status: code,
			Msg:    c.GetMsg(code),
		}
	}
	if user.CheckPassword(server.Password) == false {
		code = c.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    c.GetMsg(code),
		}
	}
	token, err := utils.GenerateToken(user.ID, server.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = c.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    c.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    c.GetMsg(code),
	}
}
