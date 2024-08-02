package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"cost"`
	PlatForm string   `json:"website"`
	Password string   `json:"-"`
	Tag      []string `json:"tags,omitempty"`
}

func main() {

	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	mycourses := []course{
		{"Go Lang Course", 900, "https://roadready.rw", "eric123", []string{"web-app", "go"}},
		{"Mern Stark Course", 350, "https://lco.com", "abcd123", []string{"Full-stark", "js"}},
		{"Laravel Course", 380, "https://lco.dev", "jkts12", nil},
	}

	// package Data in json Data
	finalJson, err := json.MarshalIndent(mycourses, "", "\t")

	if err != nil {
		panic(err)

	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	jsonFromWeb := []byte(`
	
	{
		"coursename": "Go Lang Course",
		"cost": 900,
		"website": "https://roadready.rw",
		"tags": ["web-app","go"]
     }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonFromWeb)

	if checkValid {
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)

	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// Some most cases where you want to add data to  key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}

}
