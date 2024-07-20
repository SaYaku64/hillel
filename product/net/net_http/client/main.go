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

func main() {
	resp, err := http.Get("http://localhost:8080/hello")
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

	type gg struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	ggg := gg{"name", 123}

	byteGgg, _ := json.Marshal(ggg)

	fmt.Println("Server's answer:", string(body))

	// jsonBody := []byte(`{"client_message": "hello, server!"}`)
	bodyReader := bytes.NewReader(byteGgg)

	resp, err = http.Post("http://localhost:8080/body", "application/json", bodyReader)
	if err != nil {
		fmt.Println("Post request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Post answer ReadAll error:", err)
		return
	}

	var myVar struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	json.Unmarshal(body, &myVar)

	fmt.Println("Post servers answer:", string(body))
}
