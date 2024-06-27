// ============== //
// CLIENT SERVICE //
// ============== //

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetHello() {
	resp, err := http.Get("https://sayaku2.alwaysdata.net/hello")
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Answer ReadAll error:", err)
		return
	}

	fmt.Println("Server's answer:", string(body))
}

func GetRick() {
	resp, err := http.Get("https://sayaku2.alwaysdata.net/api/v1/rick")
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Answer ReadAll error:", err)
		return
	}

	fmt.Println("Server's answer:", string(body))
}

func GetSay() {
	resp, err := http.Get("https://sayaku2.alwaysdata.net/api/v1/say?name=GOD")
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Answer ReadAll error:", err)
		return
	}
	fmt.Println("Server's answer:", string(body))
}

func GetSearchBooks() {
	type Book struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	type BookResponse struct {
		Status string `json:"status"`
		Books  []Book `json:"books"`
	}

	resp, err := http.Get("https://sayaku2.alwaysdata.net/api/v1/searchBooks?")
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Answer ReadAll error:", err)
		return
	}

	fmt.Println("Server's answer:", string(body))
}

func PostCalculate() {
	type sttt struct {
		First_num  float64 `json:"first_num"`
		Second_num float64 `json:"second_num"`
		Action     string  `json:"action"`
	}
	stt := sttt{
		First_num:  234,
		Second_num: 432,
		Action:     "addition",
	}

	sm, _ := json.Marshal(stt)
	bodyReader := bytes.NewReader(sm)
	resp, err := http.Post("https://sayaku2.alwaysdata.net/api/v1/calculate", "application/json", bodyReader)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Answer ReadAll error:", err)
		return
	}

	fmt.Println("Server's answer:", string(body))
}

func PostRegister() {
	type RegisterData struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	data := RegisterData{
		Name:     "Hesher116",
		Password: "123456778",
	}

	jsonData, _ := json.Marshal(data)
	bodyReader := bytes.NewReader(jsonData)
	resp, err := http.Post("https://sayaku2.alwaysdata.net/api/v1/register", "application/json", bodyReader)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Answer ReadAll error:", err)
		return
	}

	fmt.Println("Server's answer:", string(body))
}

func PostTranslate() {
	type TranslateData struct {
		Text           string `json:"text"`
		SourceLanguage string `json:"sourceLanguage"`
		TargetLanguage string `json:"targetLanguage"`
	}
	stt := TranslateData{
		"Hello World!",
		"en",
		"ua",
	}

	sm, _ := json.Marshal(stt)
	bodyReader := bytes.NewReader(sm)
	resp, err := http.Post("https://sayaku2.alwaysdata.net/api/v1/translate", "application/json", bodyReader)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Answer ReadAll error:", err)
		return
	}

	fmt.Println("Server's answer:", string(body))
}

// https://sayaku2.alwaysdata.net/api/v1/rick
func main() {
	//GetHello()
	//GetRick()
	//GetSay()
	//GetSearchBooks()
	//PostRegister()
	//PostCalculate()
	//PostTranslate()
}
