package main

import (
	"fmt"
	"log"
	"net/http"
)

type Resume struct {
	firstName     string `json:"firstName"`
	lastName      string `json:"lastName"`
	gender        string `json:"gender"`
	birthday      string `json:"birthday"`
	pincode       string `json:"pincode"`
	address       string `json:"address"`
	contactNumber string `json:"contactNumber"`
	emailId       string `json:"emailId"`
}

// func formValidator() {

// 	x := r.Form
// 	fmt.Println(x)
// }

func homePage(w http.ResponseWriter, r *http.Request) {

	var resumeForm []Resume

	r.ParseForm()
	fmt.Println("Endpoint Hit: homePage")

	resumeData := Resume{
		r.Form["adderrs"],
		r.Form["birthday"],
		r.Form["contact_number"],
		r.Form["email_id"],
		r.Form["firstname"],
		r.Form["lastname"],
		r.Form["pincode"],
		r.Form["gender"],
	}
	fmt.Println(resumeData)
}

func handleRequests() {
	http.HandleFunc("/submit_resume", homePage)
	log.Fatal(http.ListenAndServe(":5050", nil))
}

func main() {
	handleRequests()
}
