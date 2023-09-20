package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
} 

func main() {
    TG_API_TOKEN, exists := os.LookupEnv("TG_API_TOKEN")
    TG_API_BASE_URL := "https://api.telegram.org/bot"

    if exists {
        fmt.Println("TG token found")
    }

    resp, err := http.Get(TG_API_BASE_URL + TG_API_TOKEN + "/getMe")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(resp.Status)
    return
}
