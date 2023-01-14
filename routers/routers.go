// 负责注册应用中的路由，并引导请求到对应的控制函数中
package routers

import (
	"ginTodoList/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 初始化路由引擎
	r := gin.Default()
	// 初始化v1路由组，用于处理对应的请求
	v1Group := r.Group("v1")
	{
		// 操作待办事项
		// 添加某一个待办事项
		v1Group.POST("/todo", controller.CreateTodo)

		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	// 注册好路由信息后，返回路由引擎
	return r
}
