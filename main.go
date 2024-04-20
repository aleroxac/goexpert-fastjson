package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type Example struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	json_data := `
		{
			"string": "something", 
			"integer": 7, 
			"float": 1.2, 
			"boolean": true, 
			"array": [12.5, 6.25, 3.125], 
			"dictionary": {
				"username": "user",
				"password": "pass"
			}
		}
	`

	var parser fastjson.Parser
	parsed_data, err := parser.Parse(json_data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("string=%s\n", parsed_data.Get("string"))
	fmt.Printf("integer=%s\n", parsed_data.Get("integer"))
	fmt.Printf("float=%s\n", parsed_data.Get("float"))
	fmt.Printf("boolean=%s\n", parsed_data.Get("boolean"))
	array := parsed_data.GetArray("array")
	for i, value := range array {
		fmt.Printf("array[%d]=%s\n", i, value)
	}
	dictionary_object := parsed_data.GetObject("dictionary")
	fmt.Printf("dictionary.username=%s\n", dictionary_object.Get("username"))
	fmt.Printf("dictionary.password=%s\n", dictionary_object.Get("password"))

	var example Example
	err = json.Unmarshal([]byte(dictionary_object.String()), &example)
	if err != nil {
		panic(err)
	}
	user_info, err := json.Marshal(example)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(user_info))

	fmt.Println(parsed_data)
}
