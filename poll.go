package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-telegram/bot"
)

// createPoll creates a poll with the given question and options
func createPoll(chatID string, qn string, options []string) *bot.SendPollParams {
	poll := &bot.SendPollParams{
		ChatID:      chatID,
		Question:    qn,
		Options:     options,
		IsAnonymous: bot.False(),
	}
	return poll
}

// create daily breakfast poll
func CreateBreakfastPoll(ctx context.Context, b *bot.Bot) {
	chatID := os.Getenv("CHAT_ID")
	qn := "breakfast"
	options := []string{"8", "9", "10"}
	poll := createPoll(chatID, qn, options)
	b.SendPoll(ctx, poll)
	log.Println("Breakfast poll created at", time.Now().Format("15:04 on Monday, 2 January 2006"))
}
