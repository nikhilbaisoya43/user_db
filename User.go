package main

import "gorm.io/gorm"
 /*
 *User is used as medium to store and perfrom operation in db
 *user give json format data and decoding to the database
 */
type User struct {
	gorm.Model
	Name string `json:"name"`
	// Dob         datetime.Time `json:"dob"`
	Description string `json:"description"`
	Address     string `json:"address"`
}
