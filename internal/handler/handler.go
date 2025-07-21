package handler

import "math/rand"

var records = []model.Record{
	{ID: "1", Name: "Praneeth", AdmissionNum: rand.Intn(100)},
	{ID: "2", Name: "Sonu", AdmissionNum: rand.Intn(100)},
	{ID: "3", Name: "VR", AdmissionNum: rand.Intn(100)},
}

// write the logic for how many handlers do you want to have

func GetAllStudents() {

} // Uses GET

func GetStudentID() // Uses GET

func AddStudent() // Uses POST

func UpdateStudentInfo() // Uses PUT

func DeleteStudentRecord() // Uses DELETE
