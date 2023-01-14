// 定义结构体，读入配置文件中的参数，从而对应用中的各项服务进行配置
// 当前应用仅包含MySQL服务
package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	*MySQLConfig `ini:"mysql"`
}

// MySQLConfig 数据库配置
// 设置MySQL的用户名、密码，数据库名称、IP地址、端口号
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

// 将conf文件夹下的配置文件config.ini映射到Conf中
func Init(file string) error {
	return ini.MapTo(Conf, file)
}
