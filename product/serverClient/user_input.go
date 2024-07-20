package main

import "fmt"

type Register struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Calculate struct {
	FirstNum  float64 `json:"first_num"`
	SecondNum float64 `json:"second_num"`
	Action    string  `json:"action"`
}

type Translate struct {
	Text           string `json:"text"`
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
}

func UserInput() Register {
	var name, password string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)
	return Register{Name: name, Password: password}
}

func CalcInput() Calculate {
	var firstNum, secondNum float64
	var action string
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&firstNum)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&secondNum)
	fmt.Print("Enter the action (addition, subtraction, multiplication, division): ")
	fmt.Scanln(&action)
	return Calculate{FirstNum: firstNum, SecondNum: secondNum, Action: action}
}

func TranslateInput() Translate {
	var text, sourceLanguage, targetLanguage string
	fmt.Print("Enter the text to translate: ")
	fmt.Scanln(&text)
	fmt.Print("Enter the source language (e.g., en): ")
	fmt.Scanln(&sourceLanguage)
	fmt.Print("Enter the target language (e.g., ua): ")
	fmt.Scanln(&targetLanguage)
	return Translate{Text: text, SourceLanguage: sourceLanguage, TargetLanguage: targetLanguage}
}
