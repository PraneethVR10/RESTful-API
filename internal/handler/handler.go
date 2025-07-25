package handler

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/PraneethVR10/RESTful-API/internal/db"
	"github.com/PraneethVR10/RESTful-API/internal/model"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

var records = []model.Record{
	{ID: uuid.New().String(), Name: "Praneeth", AdmissionNum: rand.Intn(200)},
	{ID: uuid.New().String(), Name: "Sonu", AdmissionNum: rand.Intn(200)},
	{ID: uuid.New().String(), Name: "VR", AdmissionNum: rand.Intn(200)},
}

// write the logic for how many handlers do you want to have

func InsertData(c *gin.Context) {
	var insertedRecords []model.Record
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
		insertedRecords = append(insertedRecords, student)
	}
	c.IndentedJSON(http.StatusOK, insertedRecords)

}

func GetAllStudents(c *gin.Context) { // Uses GET
	rows, err := db.DB.Query(context.Background(), "SELECT id,name,admission_num FROM students")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "displaying data from the database"})
	}
	defer rows.Close()

	var data []model.Record

	for rows.Next() {
		var student model.Record
		err := rows.Scan(&student.ID, &student.Name, &student.AdmissionNum)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row", "details": err.Error()})
			return
		}
		data = append(data, student)

	}
	c.IndentedJSON(http.StatusOK, data)

}

func GetStudentID(c *gin.Context) { // Uses GET

	id := c.Param("id") // get ID from URL

	var student model.Record

	err := db.DB.QueryRow(context.Background(), "SELECT id, name, admission_num FROM students WHERE id = $1", id).Scan(&student.ID, &student.Name, &student.AdmissionNum)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found", "details": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func AddStudent(c *gin.Context) { // Uses POST
	var students []model.Record
	if err := c.BindJSON(&students); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	// Insert each student into the database
	for _, student := range students {
		_, err := db.DB.Exec(
			context.Background(),
			"INSERT INTO students (id, name, admission_num) VALUES ($1, $2, $3)",
			student.ID,
			student.Name,
			student.AdmissionNum,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to insert",
				"details": err.Error(),
			})
			return
		}
	}

	// Query all students from the database
	rows, err := db.DB.Query(context.Background(), "SELECT id, name, admission_num FROM students")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch data",
			"details": err.Error(),
		})
		return
	}
	defer rows.Close()

	// Build slice of students from result
	var allStudents []model.Record
	for rows.Next() {
		var student model.Record
		if err := rows.Scan(&student.ID, &student.Name, &student.AdmissionNum); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error scanning row",
				"details": err.Error(),
			})
			return
		}
		allStudents = append(allStudents, student)
	}

	// Return all students in JSON format
	c.IndentedJSON(http.StatusOK, allStudents)
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
