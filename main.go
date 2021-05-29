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

func storeData(resumeData Resume) {
	fmt.Println(resumeData)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println("Endpoint Hit: homePage")

	resumeData := Resume{
		address:       r.Form["address"][0],
		birthday:      r.Form["birthday"][0],
		contactNumber: r.Form["contact_number"][0],
		emailId:       r.Form["email_id"][0],
		firstName:     r.Form["firstname"][0],
		lastName:      r.Form["lastname"][0],
		pincode:       r.Form["pincode"][0],
		gender:        r.Form["gender"][0],
	}

	storeData(resumeData)
}

func handleRequests() {
	http.HandleFunc("/submit_resume", homePage)
	log.Fatal(http.ListenAndServe(":5050", nil))
}

func main() {
	handleRequests()
}
