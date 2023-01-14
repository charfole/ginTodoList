//  控制器：负责实现控制路由中对应的各个函数，其中各个函数只调用具体的逻辑，而不进行实现
/*
	url      -->  controller  -->       models
	请求来了  -->    控制器    -->   模型层的增删改查
*/
package controller

import (
	"ginTodoList/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 下面的操作主要逻辑为：接收并处理前端请求，调用models层中的增删改查基础功能进行处理
// 创建一个待办事项
func CreateTodo(c *gin.Context) {
	// 定义一个结构体用于接收前端传来的信息
	var todo models.Todo
	// 将JSON信息绑定到结构体中
	c.BindJSON(&todo)
	// 调用CreateATodo()函数，添加一个todo事项到todos表
	err := models.CreateATodo(&todo)
	// 如有错误则返回错误信息
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		// 无错误则以JSON形式返回todo信息给前端
		c.JSON(http.StatusOK, todo)
	}
}

// 获取所有待办事项
func GetTodoList(c *gin.Context) {
	// 调用模型层GetAllTodo()函数获取所有待办事项
	todoList, err := models.GetAllTodo()
	// 如有错误则返回错误信息
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		// 无错误则以JSON形式返回todo列表给前端
		c.JSON(http.StatusOK, todoList)
	}
}

// 更新一个待办事项
func UpdateATodo(c *gin.Context) {
	// 获取当前id对应的待办事项
	id, ok := c.Params.Get("id")
	// 如果没有传过来id，则返回错误
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	// 调用GetATodo()函数，首先查询出该id对应的待办事项
	todo, err := models.GetATodo(id)
	// 有错误则返回
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 将前端传来的json绑定到当前todo中
	c.BindJSON(&todo)
	// 调用UpdateATodo()函数更新todo事项
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		// 返回更新后的todo给前端（这里更新的主要是status字段，已完成或未完成）
		c.JSON(http.StatusOK, todo)
	}
}

// 删除一个待办事项
func DeleteATodo(c *gin.Context) {
	// 获取当前id对应的待办事项
	id, ok := c.Params.Get("id")
	// 如果没有传过来id，则返回错误
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	// 调用DeleteATodo()函数删除todo，删除成功则返回deleted，否则返回错误
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
