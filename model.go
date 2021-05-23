package model

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

var resume []Resume
