package handler

import (
	"math/rand"
	"net/http"

	"github.com/PraneethVR10/RESTful-API/internal/model"

	"github.com/gin-gonic/gin"
)

var records = []model.Record{
	{ID: "1", Name: "Praneeth", AdmissionNum: rand.Intn(200)},
	{ID: "2", Name: "Sonu", AdmissionNum: rand.Intn(200)},
	{ID: "3", Name: "VR", AdmissionNum: rand.Intn(200)},
}

// write the logic for how many handlers do you want to have

func GetAllStudents(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, records)

} // Uses GET

//func GetStudentID() // Uses GET

//func AddStudent() // Uses POST

//func UpdateStudentInfo() // Uses PUT

//func DeleteStudentRecord() // Uses DELETE
