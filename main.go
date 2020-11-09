package main

import (
	"fmt"
	"log"
	"strings"

	arg "github.com/alexflint/go-arg"
	"github.com/pechorka/tg"
	"github.com/pechorka/trellohelper/trello"
)

func buildReminderMsg(cards []trello.Card) string {
	var res strings.Builder
	res.WriteString("Что делать сегодня:\n\n")

	for _, card := range cards {
		if !card.Closed {
			res.WriteString(fmt.Sprintf("[%s](%s) \\- %s\n\n", card.Name, card.URL, card.Desc))
		}
	}
	return res.String()
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	var cfg struct {
		// required for work
		Key        string `arg:"env:TRELLO_KEY,required"`
		Token      string `arg:"env:TRELLO_TOKEN,required"`
		TgbotToken string `arg:"--tt, env:TELEGRAM_TOKEN,required"`
		TgChatID   int64  `arg:"--tc"`
		// optional, for certain commangs
		ListID string
	}

	if err := arg.Parse(&cfg); err != nil {
		return err
	}
	tc := trello.NewClient(cfg.Key, cfg.Token)
	tgc := tg.NewClient(cfg.TgbotToken, &tg.Options{
		ParseMod: tg.ParseModMDV2,
	})
	cards, err := tc.GetCards(cfg.ListID)
	if err != nil {
		return err
	}

	tgc.SendMsg(cfg.TgChatID, buildReminderMsg(cards))

	return nil
}
