package handler

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/PraneethVR10/RESTful-API/internal/db"
	"github.com/PraneethVR10/RESTful-API/internal/model"

	"github.com/gin-gonic/gin"
)

var records = []model.Record{
	{ID: "1", Name: "Praneeth", AdmissionNum: rand.Intn(200)},
	{ID: "2", Name: "Sonu", AdmissionNum: rand.Intn(200)},
	{ID: "3", Name: "VR", AdmissionNum: rand.Intn(200)},
}

// write the logic for how many handlers do you want to have

func InsertData(c *gin.Context) {
	for _, student := range records {
		_, err := db.DB.Exec(
			context.Background(), "INSERT INTO students (id, name, admission_num) VALUES ($1, $2, $3)",
			student.ID,
			student.Name,
			student.AdmissionNum,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert", "details": err.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Records inserted successfully"})
}

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
	// Bind the JSON from the request body to the NewStudent struct
	c.ShouldBindJSON(&newStudent)
	records = append(records, newStudent)
	c.IndentedJSON(http.StatusOK, newStudent)
}

func UpdateStudentInfo(c *gin.Context) { // Uses PUT
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

func DeleteStudentRecord(c *gin.Context) { // Uses DELETE

	id := c.Param("id") // Get ID from the request URL
	for i, student := range records {
		if student.ID == id {
			// Remove the student from the slice
			records = append(records[:i], records[i+1:]...)
		}
	}
}
