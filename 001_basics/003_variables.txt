package main

import "fmt"

var middleName string = "Kumar"
var lastName = "Completed"

func main() {
	var name1 string = "Rishav"
	var name2 = "Ravi"
	name3 := "Shivangi"
	var name4, name5 string = "Ankit", "Ankita"
	var name6, name7 = "Rahul", "Rhea"
	name8, name9 := "Apple", "Mango"

	fmt.Println(name1, name2, name3, middleName, lastName, name4, name5, name6, name7, name8, name9)
}

// Default Values of Variables :
// Numeric types : 0
// String type : ""
// Boolean type : false
// Slices, maps, structs, Pointers, functions type : nil

// Gopher symbol can only be used inside the functions.
// We need to use var keyword to declare package level variables.
