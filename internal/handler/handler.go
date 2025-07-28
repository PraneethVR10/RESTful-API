package handler

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/PraneethVR10/RESTful-API/internal/db"
	"github.com/PraneethVR10/RESTful-API/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx"

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
	names := []string{"Praneeth", "Sonu", "VR"}

	for _, name := range names {
		student := model.Record{
			ID:           uuid.New().String(), // NEW UUID EACH TIME
			Name:         name,
			AdmissionNum: rand.Intn(200),
		}

		_, err := db.DB.Exec(
			context.Background(),
			"INSERT INTO students (id, name, admission_num) VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING",
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
	// add student details from the response body
	var students []model.Record
	if err := c.BindJSON(&students); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	// Insert each student data into the database
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

	// Display the data from the result
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
	var updateDetails []model.Record
	if err := c.BindJSON(&updateDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid Entry",
			"details": err.Error(),
		})
		return
	}

	for _, student := range updateDetails {

		var existingData model.Record
		err := db.DB.QueryRow(context.Background(), "SELECT id FROM students WHERE id=$1", student.ID).Scan(&existingData.ID)

		if err == pgx.ErrNoRows {
			continue
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to locate student", "details": err.Error()})
			return

		}
		_, err = db.DB.Exec(
			context.Background(),
			"UPDATE students SET name = $1, admission_num = $2 WHERE id = $3",
			student.Name, student.AdmissionNum, student.ID,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed", "details": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Students updated successfully"})

}

func DeleteStudentRecord(c *gin.Context) { // Uses DELETE

	idParam := c.Param("id")

	var exists string
	err := db.DB.QueryRow(context.Background(), "SELECT id FROM students WHERE id=$1", idParam).Scan(&exists)

	if err == pgx.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking student existence", "details": err.Error()})
		return
	}

	_, err = db.DB.Exec(context.Background(), "DELETE FROM students WHERE id=$1", idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting student", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
