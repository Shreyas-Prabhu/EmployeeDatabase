package router

import (
	"github.com/Shreyas-Prabhu/EmployeeDatabase/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/employee/insert", controller.InsertController)
	r.GET("/employee/get/:id", controller.GetController)
	r.PUT("/employee/update/:id", controller.UpdateController)
	r.DELETE("/employee/delete/:id", controller.DeleteController)
	r.GET("/employee/getList", controller.PaginationController)

	return r
}
