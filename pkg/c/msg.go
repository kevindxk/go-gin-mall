/**
 * @Author: dexukong
 * @Description:
 * @File:  msg
 * @Version: 1.0.0
 * @Date: 2022/08/16 16:30
 */

package c

var MsgFlags = map[int]string{
	Success:             "ok",
	Error:               "fail",
	InvalidParams:       "参数错误",
	ErrorExistUser:      "用户名已经存在",
	ErrorFailEncryption: "密码加密失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	} else {
		return msg
	}
}
