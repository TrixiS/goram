package types

import (
	"strconv"
)

// Use this struct to specify a chat id or username
type ChatID struct {
	ID       int64
	Username string // In form of @username
}

func (c ChatID) MarshalJSON() ([]byte, error) {
	if c.ID != 0 {
		stringID := strconv.FormatInt(c.ID, 10)
		return []byte(stringID), nil
	}

	if c.Username != "" {
		return []byte("\"" + c.Username + "\""), nil
	}

	panic("invalid id")
}

// InlineKeyboardMarkup | ReplyKeyboardMarkup | ReplyKeyboardRemove | ForceReply
type Markup interface{}

// InputMediaAudio | InputMediaDocument | InputMediaPhoto | InputMediaVideo
type MediaGroupInputMedia interface{}
