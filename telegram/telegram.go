package telegram

import (
	"fmt"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

type teleBot struct {
	chatID int64
	bot    *tgbotapi.BotAPI
	logger *zap.SugaredLogger
}

type TeleBot interface {
	Start() error
}

func (t teleBot) Start() error {
	t.logger.Infow("starting bot")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := t.bot.GetUpdatesChan(u)
	if err != nil {
		return fmt.Errorf("failed to start bot: %w", err)
	}
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		text := fmt.Sprintf("%s %s @%s: %s",
			update.Message.From.FirstName,
			update.Message.From.LastName,
			update.Message.From.UserName,
			update.Message.Text)
		msg := tgbotapi.NewMessage(t.chatID, text)

		send, err := t.bot.Send(msg)
		if err != nil {
			t.logger.Errorf("failed to send message: %w", err)
		}
		t.logger.Debugf("message sent: %+v", send)
	}
	return nil
}

func NewBot(chatID int64, token string, logger *zap.SugaredLogger) TeleBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return &teleBot{
		chatID: chatID,
		bot:    bot,
		logger: logger.Named("teletBot"),
	}
}
