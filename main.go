package main

import "math/rand"

type Record struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	AdmissionNum string `json:"admissionNum"`
}

var record = []record{

	{ID: "1", Name: "Praneeth", AdmissionNum: rand.Intn(10)},
	{ID: "2", Name: "", AdmissionNum: rand.Intn(10)},
}

func main()
