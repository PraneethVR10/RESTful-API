package main

import (
	"github.com/PraneethVR10/RESTful-API/cmd/internal/db"
	"github.com/PraneethVR10/RESTful-API/cmd/internal/handler"
	"github.com/gin-gonic/gin"
)

func DbConnection() {

	db.ConnectDB()
	db.CreateDatabaseIfNotExists()
	db.CreateStudentsTable()

}
func main() {

	DbConnection()
	defer db.DB.Close()

	r := gin.Default()
	r.POST("/seed", handler.InsertData)
	r.GET("/students", handler.GetAllStudents)
	r.GET("/students/:id", handler.GetStudentID)
	r.POST("/students", handler.AddStudent)
	r.PUT("/students/update/", handler.UpdateStudentInfo)
	r.DELETE("/students/:id", handler.DeleteStudentRecord)
	r.Run(":3000")
}
