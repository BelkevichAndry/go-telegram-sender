package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main()  {
	println("hello")
	bot, err := tgbotapi.NewBotAPI("876041896:AAGEfSuB7-xrbANsgfEn-o9_U110J4mfjng")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	//updates, err := bot.GetUpdatesChan(u)
	msg := tgbotapi.NewMessage(618679036,"2")
	response, err := bot.Send(msg)
	log.Print(response)
}