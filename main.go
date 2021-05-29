package main

import (
	"fmt"
	"log"
	"net/http"
)

type Resume struct {
	firstName     string
	lastName      string
	gender        string
	birthday      string
	pincode       string
	address       string
	contactNumber string
	emailId       string
}

// func formValidator() {

// 	x := r.Form
// 	fmt.Println(x)
// }

func homePage(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println("Endpoint Hit: homePage")

	address := r.Form["address"]
	birthday := r.Form["birthday"]
	contact := r.Form["contact_number"]
	email := r.Form["email_id"]
	firstname := r.Form["firstname"]
	lastname := r.Form["lastname"]
	pincode := r.Form["pincode"]
	gender := r.Form["gender"]
	fmt.Println(address, birthday, contact, email, firstname, lastname, pincode, gender)

	// rawData := r.Form

	// for _, details := range rawData {
	// 	tagsList = append(tagsList, Resume{
	// 		address:       details.address,
	// 		birthday:      details.birthday,
	// 		contactNumber: details.contact,
	// 		emailId:       details.email_id,
	// 		firstName:     details.firstname,
	// 		lastName:      details.lastname,
	// 		pincode:       details.pincpde,
	// 		gender:        details.gender,
	// 	})
	// }

}

func handleRequests() {
	http.HandleFunc("/submit_resume", homePage)
	log.Fatal(http.ListenAndServe(":5050", nil))
}

func main() {
	handleRequests()
}
