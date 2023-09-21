package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
} 

type MessageBody struct {
    ChatID string `json:"chat_id"`
    Text   string `json:"text"`
}

func main() {
    TG_API_TOKEN, exists := os.LookupEnv("TG_API_TOKEN")
    TG_API_BASE_URL := "https://api.telegram.org/bot"

    if exists {
        fmt.Println("TG token found")
    }

    messageBody := MessageBody{
        ChatID: "@has_he_told_today",
        Text: "Ні, він сьогодні не розповів",
    }

    marshalled, err := json.Marshal(messageBody)
    if err != nil {
        log.Fatalf("impossible to marshal message headers: %s", err)
    }

    req, err := http.NewRequest("POST", TG_API_BASE_URL + TG_API_TOKEN + "/sendMessage", bytes.NewBuffer(marshalled))
    if err != nil {
        log.Fatalf("impossible to build request: %s", err)
    }

    req.Header.Set("Content-Type", "application/json")

    client := http.Client{Timeout: 10 * time.Second}

    res, err := client.Do(req)
    if err != nil {
        log.Fatalf("impossible to send request: %s", err)
    }
    log.Printf("status code: %d", res.StatusCode)    

    defer res.Body.Close()

    resBody, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatalf("impossible to read all body response: %s", err)
    }
    log.Printf("res body: %s", string(resBody))

    return
}
