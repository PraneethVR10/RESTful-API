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

func GetAllStudents(c *gin.Context) { // Uses GET

	c.IndentedJSON(http.StatusOK, records)

}

func GetStudentID(c *gin.Context) { // Uses GET

	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range records {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func AddStudent(c *gin.Context) { // Uses POST

	var newStudent model.Record

	// Bind the JSON payload from the request body to the NewStudent struct
	c.ShouldBindJSON(&newStudent)
	records = append(records, newStudent)
	c.IndentedJSON(http.StatusOK, newStudent)
}

func UpdateStudentInfo(c *gin.Context) {
	id := c.Param("id")
	var newDetails model.Record
	c.ShouldBindJSON(&newDetails)

	for i, s := range records {
		if s.ID == id {
			records[i] = newDetails
			c.JSON(http.StatusOK, newDetails)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "student not found"})

}

// Uses PUT

//func DeleteStudentRecord() // Uses DELETE
