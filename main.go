package main

import (
	"fmt"
	"log"
	"net/http"
)

type Resume struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Gender        string `json:"gender"`
	Birthday      string `json:"birthday"`
	Pincode       string `json:"pincode"`
	Address       string `json:"address"`
	ContactNumber string `json:"contactNumber"`
	EmailId       string `json:"emailId"`
}

func StoreData(resumeData Resume) {
	fmt.Println(resumeData)
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println("Endpoint Hit: homePage")

	ResumeData := Resume{
		Address:       r.Form["address"][0],
		Birthday:      r.Form["birthday"][0],
		ContactNumber: r.Form["contact_number"][0],
		EmailId:       r.Form["email_id"][0],
		FirstName:     r.Form["firstname"][0],
		LastName:      r.Form["lastname"][0],
		Pincode:       r.Form["pincode"][0],
		Gender:        r.Form["gender"][0],
	}

	StoreData(ResumeData)
}

func HandleRequests() {
	http.HandleFunc("/submit_resume", HomePage)
	log.Fatal(http.ListenAndServe(":5050", nil))
}

func main() {
	HandleRequests()
}
