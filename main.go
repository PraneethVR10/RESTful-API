package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/students", GetAllStudents)

	r.Run(":3000")
}
