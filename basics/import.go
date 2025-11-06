package main

import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("Welcome to importing packages in go")

	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("Response Status : ", resp.Status)
}

// Renaming the imported packages as shown using foo example
