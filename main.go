package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	InfinitySend()
}

func InfinitySend() {
	stop := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Second)
		stop <- struct{}{}
	}()

	for {
		select {
		case <-stop:
			log.Println("Send More Message after 10 seconds")
			time.Sleep(10 * time.Second)
			InfinitySend()
			return
		default:
			for j := 0; j < 20; j++ {
				go func() {
					_, err := sendTelegramMessage()
					if err != nil {
						log.Println(string(err.Error()))
						stop <- struct{}{}
					}
				}()
			}
			time.Sleep(1 * time.Second)
		}
	}

}

func sendTelegramMessage() (map[string]interface{}, error) {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Println("TELEGRAM_BOT_TOKEN is not set")
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN is not set")
	}

	TELEGRAM_CHAT_ID := os.Getenv("TELEGRAM_CHAT_ID")
	if TELEGRAM_CHAT_ID == "" {
		log.Println("TELEGRAM_CHAT_ID is not set")
		return nil, fmt.Errorf("TELEGRAM_CHAT_ID is not set")
	}

	chatid := fmt.Sprintf("%s", TELEGRAM_CHAT_ID)

	message := "\nBerhasil Kirim SMS dari jauh \nKepada : "
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(99999999)
	message += fmt.Sprintf("%08d", randomNumber)
	pesan := "Setiap kesulitan adalah peluang untuk tumbuh dan berkembang. Terimalah dengan tabah, karena di baliknya tersimpan hikmah yang berharga."
	message += fmt.Sprintf(" \nPesan : %s", pesan)

	escapedMessage := url.QueryEscape(message)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?parse_mode=Markdown&chat_id="+chatid+"&text=%s", token, escapedMessage)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to send message to telegram: %v", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	log.Println("Message sent to telegram url: ", url)
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
