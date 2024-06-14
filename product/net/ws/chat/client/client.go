package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run client.go <username>")
	}
	username := os.Args[1]
	url := "ws://localhost:8080/ws?username=" + username
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}
			var msg Message
			if err := json.Unmarshal(message, &msg); err != nil {
				log.Println("Unmarshal error:", err)
				continue
			}
			log.Printf("[%s]: %s", msg.Username, msg.Text)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			msg := Message{
				Username: username,
				Text:     text,
			}
			msgBytes, err := json.Marshal(msg)
			if err != nil {
				log.Println("Marshal error:", err)
				return
			}
			if err := conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
				log.Println("Write error:", err)
				return
			}
		}
	}()

	<-interrupt
	log.Println("Interrupt received, closing connection...")
	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	time.Sleep(time.Second)
}
