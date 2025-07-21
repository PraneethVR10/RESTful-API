package main

import (
	"github.com/PraneethVR10/RESTful-API/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/students", handler.GetAllStudents)
	r.GET("/students/:id", handler.GetStudentID)

	r.Run(":3000")
}
