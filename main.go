package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Company struct {
	Name        string
	CatchPhrase string
	BS          string
}

type User struct {
	Name    string
	Email   string
	Company Company
}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return
	}
	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON", err)
	}

	for idx, user := range users {
		list := idx + 1
		fmt.Printf("---User %d---\n\n", list)
		fmt.Printf("Name: %s\n", user.Name)
		fmt.Printf("Email: %s\n", user.Email)
		fmt.Println("Company:")
		fmt.Printf("  Name: %s\n", user.Company.Name)
		fmt.Printf("  Catch Phrase: %s\n", user.Company.CatchPhrase)
		fmt.Printf("  Business Service: %s\n", user.Company.BS)
		fmt.Println("")
	}
}
