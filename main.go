package main

import (
    "fmt"
	"bytes"
	"encoding/json"
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

func sendBasicMessage(endpoint string, messageText string) {
    messageBody := MessageBody{
        ChatID: "@has_he_told_today",
        Text: messageText,
    }

    marshalled, err := json.Marshal(messageBody)
    if err != nil {
        log.Fatalf("impossible to marshal message headers: %s", err)
    }

    req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(marshalled))
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
}

func startPolling(endpoint string) {
    for {
        time.Sleep(5 * time.Second)
        go sendBasicMessage(endpoint, "Ні, він сьогодні не розповів")
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    TG_API_TOKEN, _ := os.LookupEnv("TG_API_TOKEN")
    TG_API_BASE_URL := "https://api.telegram.org/bot"

    basicEndpoint := TG_API_BASE_URL +  TG_API_TOKEN + "/sendMessage"
    go startPolling(basicEndpoint)

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
