package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string  `json:"first_name"`          // FirstName in struct is mapped to first_name in json format.
	LastName  string  `json:"last_name,omitempty"` // omitempty will prevent the occurrance of last_name in json format if no value is provided to LastName in struct.
	Age       int     `json:"age"`
	Address   Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Employee struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Gender  string  `json:"gender,omitempty"`
	Address Address `json:"address"`
}

func main() {
	person := Person{ // note that LastName and Address.State values are not given to the struct. But LastName will only disappear from the json format as omitempty tag is included there but Address.State will take the default value as "" and appear there.
		FirstName: "Rishav",
		Age:       26, // similarly if Age is also not provided any value in the struct, it will assume default int value as 0 and appear in the json format.
		Address: Address{
			City: "Jamshedpur",
		},
	}

	jsonData1, err := json.Marshal(person) // marshall is used to convert struct into json. Here jsonData1 is []byte slice.
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println("Person :", string(jsonData1)) // json format will be {"first_name":"Rishav","age":26,"address":{"city":"Jamshedpur","state":""}}

	jsonData2 := `{"name" : "Rishav Kumar", "age" : 30, "address" : {"state" : "Jharkhand"}}` // sample json format data

	var employee Employee
	err = json.Unmarshal([]byte(jsonData2), &employee) // unmarshall is used to convert the json into struct. Here struct must be passed as reference.
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println("Employee :", employee)
	fmt.Println("Employee State :", employee.Address.State) // any field of the struct can be properly handled now.

	// _____________________Marshaling/Unmarshaling Lists____________________________________________________________________________________________________________________________________

	addressList1 := []Address{
		{City: "Jamshedpur", State: "Jharkhand"},
		{City: "New Delhi"},
		{City: "Mumbai", State: "Maharastra"},
	}
	jsonList1, err := json.Marshal(addressList1) // converts slice of Address struct to json.
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println("Addresses :", string(jsonList1)) // prints [{"city":"Jamshedpur","state":"Jharkhand"},{"city":"New Delhi","state":""},{"city":"Mumbai","state":"Maharastra"}]

	jsonList2 := `[{"city": "New York", "state": "NY"}, {"city": "Sydney"}, {"state": "Jharkhand"}]`
	var addressList2 []Address
	err = json.Unmarshal([]byte(jsonList2), &addressList2) // converts json to slice of Address struct
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println("Addresses :", addressList2)                   // prints [{New York NY} {Sydney } { Jharkhand}]
	fmt.Println("City of 2nd address :", addressList2[1].City) // prints Sydney

	personList := []Person{
		{FirstName: "Rishav", LastName: "Kumar", Age: 26, Address: Address{City: "Jamshedpur", State: "Jharkhand"}},
		{FirstName: "Saurav", LastName: "Kumar", Age: 26},
		{FirstName: "Ashish", Address: Address{State: "Jharkhand"}},
	}
	jsonList3, err := json.Marshal(personList) // converts the slice of Person struct to json.
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println("People :", string(jsonList3)) // prints the output as [{"first_name":"Rishav","last_name":"Kumar","age":26,"address":{"city":"Jamshedpur","state":"Jharkhand"}},{"first_name":"Saurav","last_name":"Kumar","age":26,"address":{"city":"","state":""}},{"first_name":"Ashish","age":0,"address":{"city":"","state":"Jharkhand"}}]

	jsonList4 := `[{"name": "Rishav Kumar", "age": 30, "gender": "male", "address": {"city": "Jamshedpur", "state": "Jharkhand"}}, {"name": "Akash Yadav", "gender": "male"}, {"name": "Ravi Kumar", "address": {"state": "Bihar"}}]`
	var employeeList1 []Employee                            // using proper struct slice to store the data if the structure of the json data is known.
	err = json.Unmarshal([]byte(jsonList4), &employeeList1) // converts json to slice of Employee struct.
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println("EmployeeList1 :", employeeList1) // prints the output as [{Rishav Kumar 30 male {Jamshedpur Jharkhand}} {Akash Yadav 0 male { }} {Ravi Kumar 0  { Bihar}}]

	var employeeList2 []map[string]interface{}              // using slice of map with key as string and value as interface{} if structure of json data is not known.
	err = json.Unmarshal([]byte(jsonList4), &employeeList2) // converts json to slice of map[string]interface{}
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println("EmployeeList2 :", employeeList2) // prints the output as [map[address:map[city:Jamshedpur state:Jharkhand] age:30 gender:male name:Rishav Kumar] map[gender:male name:Akash Yadav] map[address:map[state:Bihar] name:Ravi Kumar]]
}
