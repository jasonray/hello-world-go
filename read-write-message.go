package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Message struct {
	ID      string `json:"id"`
	Payload string `json:"payload"`
}

func main() {
	// remove and create directory to start clean
	err := os.RemoveAll("/tmp/hello-world-go")
	check(err)
	err = os.Mkdir("/tmp/hello-world-go/", 0755)
	check(err)

	// write some static text to a file
	output_string := "hello world!"
	boutput := []byte(output_string)
	file_path := "/tmp/hello-world-go/o1.txt"
	err2 := os.WriteFile(file_path, boutput, 0644)
	check(err2)

	// write some data from a struct to file (manual txt format)
	m1 := Message{ID: "1", Payload: "this is first message"}
	fmt.Println("m1:", m1)
	m1_txt := "id: " + m1.ID + "\n" + "payload: " + m1.Payload
	file_path = "/tmp/hello-world-go/m1.txt"
	err = os.WriteFile(file_path, []byte(m1_txt), 0644)
	check(err)

	// write some data from a struct to file (json format)
	m1_json, err := json.Marshal(m1)
	// print("m1_json: ", m1_json)
	check(err)
	file_path = "/tmp/hello-world-go/m1.json"
	err = os.WriteFile(file_path, m1_json, 0644)
	check(err)

	m2 := Message{ID: "2", Payload: "this is second message"}
	fmt.Println("m2:", m2)
	m2_json, err := json.Marshal(m2)
	// fmt.Println("m2_json: ", m2_json)
	check(err)
	file_path_m2 := "/tmp/hello-world-go/m2.json"
	err = os.WriteFile(file_path_m2, m2_json, 0644)
	check(err)

	m3_json, err := os.ReadFile(file_path_m2)
	check(err)
	m3 := Message{}
	err = json.Unmarshal(m3_json, &m3)
	check(err)
	fmt.Println("m3:", m3)

	m3_json, err = json.MarshalIndent(m3, "", "   ")
	file_path_m3 := "/tmp/hello-world-go/m3.json"
	err = os.WriteFile(file_path_m3, m3_json, 0644)
	check(err)

}

// func main() {
// 	// Create an instance of the Message struct
// 	message := Message{
// 		id:      "123",
// 		ID:      "345",
// 		payload: "Hello, World!",
// 		Payload: "hola",
// 	}

// 	// Marshal the struct into JSON format
// 	jsonData, err := json.Marshal(message)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON:", err)
// 		return
// 	}

// 	// Open a file for writing
// 	file, err := os.Create("/tmp/message.json")
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Write the JSON data to the file
// 	_, err = file.Write(jsonData)
// 	if err != nil {
// 		fmt.Println("Error writing JSON data to file:", err)
// 		return
// 	}

// 	fmt.Println("Message written to message.json")
// }
