/*/**
 * @Author: dexukong
 * @Description:
 * @File:  code
 * @Version: 1.0.0
 * @Date: 2022/08/16 16:29
 */

package e

const (
	Success             = 200
	Error               = 500
	InvalidParams       = 400
	ErrorExistUser      = 30001
	ErrorFailEncryption = 30002
	ErrorUserNotFond    = 30003
	ErrorNotCompare     = 30004
	ErrorAuthToken      = 30005

	//管理员错误

	ErrorAuthCheckTokenFail        = 30001 //token 错误
	ErrorAuthCheckTokenTimeout     = 30002 //token 过期
	ErrorAuth                      = 30004
	ErrorAuthInsufficientAuthority = 30005
	ErrorReadFile                  = 30006
	ErrorSendEmail                 = 30007
	ErrorCallApi                   = 30008
	ErrorUnmarshalJson             = 30009
	ErrorAdminFindUser             = 30010
	//数据库错误
	ErrorDatabase = 40001
)
