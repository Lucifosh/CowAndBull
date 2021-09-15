package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func TelegramBot() {
	curGame := make(map[int]Game)
	bot, err := tgbotapi.NewBotAPI(GetToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	for {
		updates, _ := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message == nil {
				continue
			}
			if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

				text := update.Message.Text
				switch {
				case strings.Contains(text, "/start") || strings.Contains(strings.ToLower(text), "старт"):
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я игровой бот. Чтобы сыграть со мной введи \"играть\".")
					bot.Send(msg)
				case strings.Contains(strings.ToLower(text), "играть"):
					g := NewBullAndCow()
					g.state = "go"
					start := time.Now()
					g.time = start
					curGame[update.Message.From.ID] = g
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Игра началась!\n")
					bot.Send(msg)
				case strings.Contains(strings.ToLower(text), "рейтинг"):
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, SendBoard(update.Message.From.ID))
					bot.Send(msg)
				case strings.Contains(strings.ToLower(text), "help") || strings.Contains(strings.ToLower(text), "помощь") || strings.Contains(strings.ToLower(text), "правила"):
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваша задача отгадать 4х значное число.\nКоровы - Вы отгадали одно из чисел, но не угадали с позицией.\nБыки - Вы отгадали число и его позицию.\nУдачи!")
					bot.Send(msg)
				default:
					g := curGame[update.Message.From.ID]
					if g.state == "go" {
						cow, bull, err := CheckNumber(g.number, text)
						if err != "" {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, err)
							bot.Send(msg)
						} else if bull == 4 {
							g.state = "win"
							t := time.Since(g.time).Seconds()
							score := 0
							if t < 60 {
								score = 10
							} else if t < 90 {
								score = 7
							} else {
								score = 5
							}

							m1 := fmt.Sprintf("Поздравляю!\nХодов: %v\nВремя: %.2f секунд\n", g.count, t)
							tempScore := getScoreFromID(update.Message.From.ID)
							m2 := fmt.Sprintf("Рейтинг: %v + %v -> %v", score, tempScore, score+tempScore)
							WriteBoard(update.Message.From.ID, score)
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, m1+m2)
							bot.Send(msg)
						} else {
							g.count = g.count + 1
							m := fmt.Sprintf("коровы: %v\nбыки:   %v\n", cow, bull)
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, m)
							bot.Send(msg)
						}
						curGame[update.Message.From.ID] = g
					} else {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Попробуйте ввести одну из следующих команд без кавычек:\n\"играть\"\n\"помощь\"\n\"рейтинг\"")
						bot.Send(msg)
					}
				}
			}
		}
	}
}
