/**
 * @Author: dexukong
 * @Description:
 * @File:  msg
 * @Version: 1.0.0
 * @Date: 2022/08/16 16:30
 */

package e

var MsgFlags = map[int]string{
	Success:             "ok",
	Error:               "fail",
	InvalidParams:       "参数错误",
	ErrorExistUser:      "用户名已经存在",
	ErrorFailEncryption: "密码加密失败",
	ErrorUserNotFond:    "用户不存在",
	ErrorNotCompare:     "密码错误",
	ErrorAuthToken:      "Token错误",
	ErrorDatabase:       "数据库错误",

	////
	//ErrorAuthCheckTokenFail:    "Token鉴权失败",
	//ErrorAuthCheckTokenTimeout: "Token已超时",
	//ErrorAuthToken:                 "Token生成失败",
	//ErrorAuth:                      "Token错误",
	//ErrorAuthInsufficientAuthority: "权限不足",
	ErrorReadFile:      "读文件失败",
	ErrorSendEmail:     "发送邮件失败",
	ErrorCallApi:       "调用接口失败",
	ErrorUnmarshalJson: "解码JSON失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	} else {
		return msg
	}
}
