package controllers

import (
	"CrudORM/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (c *Controller) CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	ctx.BindJSON(&todo)
	c.DB.Create(&todo)
	ctx.JSON(200, todo)
}

func (c *Controller) GetTodos(ctx *gin.Context) {
	var todos []models.Todo
	c.DB.Find(&todos)
	ctx.JSON(200, todos)
}

func (c *Controller) GetTodo(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var todo models.Todo
	c.DB.First(&todo, id)
	ctx.JSON(200, todo)
}

func (c *Controller) UpdateTodo(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var todo models.Todo
	c.DB.First(&todo, id)
	ctx.BindJSON(&todo)
	c.DB.Save(&todo)
	ctx.JSON(200, todo)
}

func (c *Controller) DeleteTodo(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	fmt.Println(id)
	var todo models.Todo
	c.DB.Where("id = ?", id).Delete(&todo)
	ctx.JSON(200, gin.H{"id #" + id: "deleted"})
}
