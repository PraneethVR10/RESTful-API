package main

import (
	"github.com/PraneethVR10/RESTful-API/internal/db"
	"github.com/PraneethVR10/RESTful-API/internal/handler"
	"github.com/gin-gonic/gin"
)

func DbConnection() {

	db.ConnectDB()
	defer db.DB.Close()
	db.CreateDatabaseIfNotExists()
}
func main() {

	DbConnection()

	r := gin.Default()
	r.GET("/students", handler.GetAllStudents)
	r.GET("/students/:id", handler.GetStudentID)
	r.POST("/students", handler.AddStudent)
	r.PUT("/students/update/:id", handler.UpdateStudentInfo)
	r.DELETE("/students/:id", handler.DeleteStudentRecord)
	r.Run(":3000")
}
