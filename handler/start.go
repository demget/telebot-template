package handler

import (
	tele "gopkg.in/tucnak/telebot.v3"
)

func (h Handler) OnStart(c tele.Context) error {
	return c.Send(h.Lt(c).Text("start"))
}
