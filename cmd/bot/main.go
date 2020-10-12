package main

import (
	"log"
	"os"

	tele "gopkg.in/tucnak/telebot.v3"
	"gopkg.in/tucnak/telebot.v3/layout"
)

func main() {
	lt, err := layout.New("bot.yml")
	if err != nil {
		log.Fatal(err)
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.Open(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	h := handler.New(handler.Config{
		Layout: lt,
		Bot:    b,
		DB:     db,
	})

	// Middleware
	b.OnError = h.OnError
	b.Use(lt.Middleware("en", h.LocaleFunc))

	// Handlers
	b.Handle("/start", h.OnStart)

	b.Start()
}
