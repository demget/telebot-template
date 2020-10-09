package handler

import (
	tele   "gopkg.in/tucnak/telebot.v3"
	telelt "gopkg.in/tucnak/telebot.v3/layout"
)

func New(c Handler) handler {
	return handler{
		lt: c.Layout,
		b:  c.Bot,
		db: c.DB,
	}
}

type (
	Handler struct {
		Layout *telelt.Layout
		Bot    *tele.Bot
		DB     *storage.DB
	}
	handler struct {
		lt *telelt.Layout
		b  *tele.Bot
		db *storage.DB
	}
)
