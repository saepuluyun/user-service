package models

import ()

// User models is model that defening user
type User struct {
	UID       string `json:"uid"`
	FirstName string `json:"first_name`
	LastName  string `json:"last_name"`
}
