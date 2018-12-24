package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/robfig/cron"
)

func main() {
	bot := connectBot()

	c := cron.New()
	// Send Message Everyday @ 06:00 AM
	c.AddFunc("0 0 6 * * *", func() {
		sendGreetingText(bot, "{USER_ID|GROUP_ID|ROOM_ID}")
	})
	c.Start()

	staticFileServer := http.FileServer(http.Dir("static"))
	http.HandleFunc("/static/", http.StripPrefix("/static/", staticFileServer).ServeHTTP)
	http.HandleFunc("/callback", callBackHandler(bot))

	port := getPort()
	log.Printf("Serving on localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func connectBot() *linebot.Client {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	return bot
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3030"
	}

	return ":" + port
}

func callBackHandler(bot *linebot.Client) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					currentDate := getCurrentThaiDate()
					log.Printf("New Message: %s\n", message.Text)
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
					bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(currentDate)).Do()
				}
			}
		}
	}
}

func sendGreetingText(bot *linebot.Client, to string) {
	now := getCurrentThaiDate()
	weekday := fmt.Sprint(time.Now().Weekday())
	bot.PushMessage(to, linebot.NewTextMessage(now)).Do()

	appURL := os.Getenv("APP_URL")
	imagePath := appURL + "/static/" + weekday + ".jpg"

	log.Print("Send Messages")
	bot.PushMessage(to, linebot.NewImageMessage(imagePath, imagePath)).Do()
}
