package markups

import (
	"github.com/TrixiS/goram"
	"github.com/TrixiS/goram/cbdata"
	"github.com/TrixiS/goram/keyboards"
)

const Prefix = "start"

type CbDataType uint8

const (
	Hello CbDataType = iota
	World
)

type CbData struct {
	Type CbDataType
}

var Start = &goram.InlineKeyboardMarkup{
	InlineKeyboard: keyboards.NewBuilder[goram.InlineKeyboardButton]().Row(
		goram.InlineKeyboardButton{
			Text:         "Hello",
			CallbackData: cbdata.Pack(Prefix, CbData{Type: Hello}),
		},
		goram.InlineKeyboardButton{
			Text:         "World",
			CallbackData: cbdata.Pack(Prefix, CbData{Type: World}),
		},
	).Adjust(1).Build(),
}
