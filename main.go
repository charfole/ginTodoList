// 整个应用的入口
package main

import (
	"fmt"
	"ginTodoList/dao"
	"ginTodoList/models"
	"ginTodoList/routers"
	"ginTodoList/setting"
	"os"
)

func main() {
	// 1. 首先需要创建数据库，在MySQL中输入以下命令
	// sql: CREATE DATABASE ginTodoList DEFAULT CHARSET=utf8mb4;

	// 2. 以 ./ginTodoList conf/config.ini 的形式启动项目
	if len(os.Args) != 2 {
		fmt.Println("Usage：./ginTodoList conf/config.ini")
		return
	}
	// 3. 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	// 4. 连接数据库
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	// 5. 程序退出的时候关闭数据库连接
	defer dao.Close()

	// 6. 模型绑定，将Todo结构体绑定到todos表中
	dao.DB.AutoMigrate(&models.Todo{})

	// 7. 注册路由以响应前端的需求，注册完成后启动服务
	r := routers.SetupRouter()
	r.Run(":9090")
}
