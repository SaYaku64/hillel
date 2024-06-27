package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	getUrls := []string{
		"https://sayaku2.alwaysdata.net/hello",
		"https://sayaku2.alwaysdata.net/api/v1/rick",
		"https://sayaku2.alwaysdata.net/api/v1/say?name=Max",
		"https://sayaku2.alwaysdata.net/api/v1/searchBooks?title=Harry%20Potter&author=Joanne%20Rowling",
	}
	postUrls := []string{
		"https://sayaku2.alwaysdata.net/api/v1/register",
		"https://sayaku2.alwaysdata.net/api/v1/calculate",
		"https://sayaku2.alwaysdata.net/api/v1/translateText",
	}

	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Perform GET requests")
		fmt.Println("2. Perform POST request to /register")
		fmt.Println("3. Perform POST request to /calculate")
		fmt.Println("4. Perform POST request to /translateText")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			for _, url := range getUrls {
				body, err := fetchGETURL(url)
				if err != nil {
					log.Println(err)
					continue
				}
				fmt.Printf("Server's answer from %s: %s\n", url, body)
			}
		case 2:
			user := UserInput()
			data, err := json.Marshal(user)
			if err != nil {
				log.Fatalf("Error marshalling data: %v", err)
			}
			response, err := fetchPOSTURL(postUrls[0], data)
			if err != nil {
				log.Fatalf("Error making POST request: %v", err)
			}
			fmt.Printf("Server's answer: %s\n", response)
		case 3:
			calcInput := CalcInput()
			data, err := json.Marshal(calcInput)
			if err != nil {
				log.Fatalf("Error marshalling data: %v", err)
			}
			response, err := fetchPOSTURL(postUrls[1], data)
			if err != nil {
				log.Fatalf("Error making POST request: %v", err)
			}
			fmt.Printf("Server's answer: %s\n", response)
		case 4:
			translateInput := TranslateInput()
			data, err := json.Marshal(translateInput)
			if err != nil {
				log.Fatalf("Error marshalling data: %v", err)
			}
			response, err := fetchPOSTURL(postUrls[2], data)
			if err != nil {
				log.Fatalf("Error making POST request: %v", err)
			}
			fmt.Printf("Server's answer: %s\n", response)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
