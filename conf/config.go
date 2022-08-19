/**
 * @Author: dexukong
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/08/15 14:33
 */

package conf

import (
	"ginmall/dao"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppModel string
	HttpPort string

	DB         string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatorPath  string
)

func Init() {
	//本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMysql(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPath(file)

	//mysql读
	pathRead := strings.Join([]string{DBUser, ":", DBPassword, "@tcp(", DBHost, ":", DBPort, ")/", DBName, "?charset=utf8mb4&parseTime=true"}, "")
	//mysql写的
	pathWrite := strings.Join([]string{DBUser, ":", DBPassword, "@tcp(", DBHost, ":", DBPort, ")/", DBName, "?charset=utf8mb&parseTime=true"}, "")

	dao.Database(pathRead, pathWrite)
}

func LoadServer(file *ini.File) {
	AppModel = file.Section("service").Key("AppModel").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DBHost = file.Section("mysql").Key("DBHost").String()
	DBPort = file.Section("mysql").Key("DBPort").String()
	DBUser = file.Section("mysql").Key("DBUser").String()
	DBPassword = file.Section("mysql").Key("DBPassword").String()
	DBName = file.Section("mysql").Key("DBName").String()

}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}

func LoadPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatorPath = file.Section("path").Key("AvatorPath").String()
}
