package main

import "fmt"

func main() {
	name := "Rishav"
	age := 26

	//____________Printing Functions___________________________________________________________________________________________
	fmt.Print("Hello", "World!", 123, 456)         // prints without breaking into newline.
	fmt.Println("Hello", "World!", 123, 456)       // prints and breaks into the newline.
	fmt.Printf("Name : %s, Age : %d\n", name, age) // prints according to the format provided.

	//____________Formatting Functions_________________________________________________________________________________________
	s1 := fmt.Sprint("Hello", "World!", 123, 456)         // returns the string without the newline at the end.
	s2 := fmt.Sprintln("Hello", "World!", 123, 456)       // returns the string with the newline at the end.
	s3 := fmt.Sprintf("Name : %s, Age : %d\n", name, age) // returns the string in the asked format.
	fmt.Print(s1)
	fmt.Print(s2)
	fmt.Print(s3)

	//____________Scanning Functions___________________________________________________________________________________________
	var username1, username2, username3 string
	var password1, password2, password3 int
	fmt.Print("Enter the first username(string) and password(int) : ")
	fmt.Scan(&username1, &password1) // waits for entering all the input values irrespective of whether you go to newline or not.

	fmt.Print("Enter the second username(string) and password(int) : ")
	fmt.Scanln(&username2, &password2) // entering the newline means all inputs are taken. If no value is given to input then the zero value of the corresponding varaible type is taken as input.

	fmt.Print("Enter the third username(string) and password(int) : ")
	fmt.Scanf("%s , %d", &username3, &password3) // expects the inputs in the provided format only. Example - rishavkumar , 123456

	fmt.Printf("Username1 : %s, Password1 : %d\n", username1, password1)
	fmt.Printf("Username2 : %s, Password2 : %d\n", username2, password2)
	fmt.Printf("Username3 : %s, Password3 : %d\n", username3, password3)

	//____________Error Function_______________________________________________________________________________________________
	err := checker(-2)
	if err != nil {
		fmt.Println("Error :", err)
	}
}

func checker(val int) error {
	if val < 0 {
		return fmt.Errorf("Value can't be negative.") // returns the error string in the specified format. It also satisfies the error interface.
	}
	return nil
}

// fmt.Print(), fmt.Sprint() functions add an extra space only if none of those two adjacent operands are strings otherwise not.
// fmt.Println(), fmt.Sprintln() functions add an extra space between every two operands irrespective of strings or not.
