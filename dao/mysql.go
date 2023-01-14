// dao(Data Access Object)：负责初始化数据库连接与关闭数据库连接
package dao

import (
	"fmt"
	"ginTodoList/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义数据库对象
var (
	DB *gorm.DB
)

// 初始化MySQL数据库，指定数据库用户、密码、IP、端口号和数据库名称信息
func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

// 关闭数据库
func Close() {
	DB.Close()
}
