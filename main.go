package main

import (
	"math/rand"
	"time"

	"github.com/PraneethVR10/RESTful-API/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := gin.Default()
	r.GET("/students", handler.GetAllStudents)

	r.Run(":3000")
}
