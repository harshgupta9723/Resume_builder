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

func formValidator(resume Resume) {

}

func homePage(w http.ResponseWriter, r *http.Request) {

	var resumeForm []Resume

	r.ParseForm()
	fmt.Println("Endpoint Hit: homePage")

	for _, value := range r.Form {
		resumeForm = append(resumeForm, Resume{
			firstName:     value.firstname,
			lastName:      value.lastname,
			gender:        value.gender,
			address:       value.address,
			contactNumber: value.contact_number,
			emailId:       value.email_id,
			birthday:      value.birthday,
			pincode:       value.pincode,
		})
	}
	fmt.Println(resumeForm)

}

func handleRequests() {
	http.HandleFunc("/submit_resume", homePage)
	log.Fatal(http.ListenAndServe(":5050", nil))
}

func main() {
	handleRequests()
}
