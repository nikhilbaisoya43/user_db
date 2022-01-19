package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	// Dob         datetime.Time `json:"dob"`
	Description string `json:"description"`
	Address     string `json:"address"`
}
