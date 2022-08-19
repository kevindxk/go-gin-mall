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
	"ginmall/conf"
	"ginmall/dao"
	"ginmall/model"
	"ginmall/pkg/e"
	"ginmall/serializer"
	"ginmall/utils"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/mail.v2"
	"strings"
	"time"
)

type UserServer struct {
	NickName string `json:"nick_name" from:"nick_name"  `
	UserName string `json:"user_name" from:"user_name"`
	Password string `json:"password" from:"password"`
	Key      string `json:"nick_name" from:"key"`
}

type SendEmailService struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OpertionType uint `from:"opertion_type" json:"opertion_type"`
}

type ValidEmailService struct {
}

//用户注册
func (server *UserServer) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if server.Key == "" || len(server.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}
	//---->密文加密存储
	utils.Encrypt.SetKey(server.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(server.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
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
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//用户登录
func (server *UserServer) Login(ctx context.Context) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(server.UserName)
	if !exist {
		logging.Info(err)
		code = e.ErrorUserNotFond
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if user.CheckPassword(server.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := utils.GenerateToken(user.ID, server.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

//用户更新
func (server *UserServer) Update(ctx context.Context, uid uint) serializer.Response {
	var user model.User
	code := e.Success
	//找到用户
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if server.NickName != "" {
		user.NickName = server.NickName
	}
	err = userDao.UpdateUserById(uid, user)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(&user),
		Msg:    e.GetMsg(code),
	}
}

func (sendEmail *SendEmailService) Send(ctx context.Context, id uint) serializer.Response {
	code := e.Success
	var address string
	token, err := utils.GenerateEmailToken(id, sendEmail.OpertionType, sendEmail.Email, sendEmail.Password)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	userDao := dao.NewNoticeDao(ctx)
	notice, err := userDao.GetNoticeById(sendEmail.OpertionType)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = conf.ValidEmail + token
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", sendEmail.Email)
	m.SetHeader("Subject", "FanOne")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		logging.Info(err)
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Valid 验证内容
func (service ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userID uint
	var email string
	var password string
	var operationType uint
	code := e.Success

	//验证token
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := utils.ParseEmailToken(token)
		if err != nil {
			logging.Info(err)
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userID = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//获取该用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userID)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if operationType == 1 {
		//1:绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		//2：解绑邮箱
		user.Email = ""
	} else if operationType == 3 {
		//3：修改密码
		err = user.SetPassword(password)
		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserById(userID, user)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 成功则返回用户的信息
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(&user),
	}
}
