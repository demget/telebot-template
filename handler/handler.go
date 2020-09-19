package handler

import (
	tele   "gopkg.in/tucnak/telebot.v3"
	telelt "gopkg.in/tucnak/telebot.v3/layout"
)

func New(c Config) Handler {
	return Handler{
		lt: c.Layout,
		b:  c.Bot,
		db: c.DB,
	}
}

type (
	Handler struct {
		lt *telelt.Layout
		b  *tele.Bot
		db *storage.DB
	}
	Config struct {
		Layout *telelt.Layout
		Bot    *tele.Bot
		DB     *storage.DB
	}
)

func (h Handler) Lt(c tele.Context) *telelt.Layout {
	return h.lt.With(c)
}