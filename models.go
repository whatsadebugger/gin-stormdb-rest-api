package main

import ()

type address struct {
	ID        int
	FirstName string `json:"firstname" storm:"id,unique"`
	LastName  string `json:"lastname"  storm:"index"`
	Email     string `json:"email" 		 storm:"index,unique"`
	Phone     string `json:"phone"		 storm:"index,unique"`
}

var addressBook []address
