package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
	"github.com/mileusna/crontab"
)

func main() {
	// environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// create context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// create bot
	b, err := bot.New(os.Getenv("BOT_TOKEN"), bot.WithDefaultHandler(handler))
	if err != nil {
		log.Fatal(err)
	}

	// repeat job at 10pm everyday
	ctab := crontab.New()
	ctab.MustAddJob("0 22 * * 0-5", func() { CreateBreakfastPoll(ctx, b) })

	// start bot
	log.Println("Bot started")
	b.Start(ctx)
	log.Println("Bot shut down")
}

func handler(_ context.Context, _ *bot.Bot, _ *models.Update) {
	// ignore all updates
	return
}
