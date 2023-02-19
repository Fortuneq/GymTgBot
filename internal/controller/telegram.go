package controller

import (
	tele "gopkg.in/telebot.v3"
)

type GymBot struct {
	Bot *tele.Bot
}

type Bot interface {
	Start() error
	startEndpointHandler(ctx tele.Context) error
}

func (b *GymBot) Start() error {
	b.Bot.Handle("/start", b.startEndpointHandler)

	return nil
}

func (b *GymBot) startEndpointHandler(c tele.Context) error {
	msg := c.Message()
	c.Send(msg)
	return nil
}

func NewBotController(bot *tele.Bot) *GymBot {
	return &GymBot{bot}
}

var _ Bot = &GymBot{}
