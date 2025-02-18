package types

import (
	"io"
	"strconv"
)

// Use this struct to specify a chat id or username
type ChatID struct {
	ID       int64  // Negative int64 (-100...) for channel and some group ids, positive for user ids
	Username string // Plain username, without leading @
}

// For json encoding ChatID inside structs
func (c ChatID) MarshalJSON() ([]byte, error) {
	if c.ID != 0 {
		stringID := strconv.FormatInt(c.ID, 10)
		return []byte(stringID), nil
	}

	if c.Username != "" {
		return []byte("\"" + "@" + c.Username + "\""), nil
	}

	return nil, nil
}

func (c ChatID) String() string {
	if c.ID != 0 {
		stringID := strconv.FormatInt(c.ID, 10)
		return stringID
	}

	if c.Username != "" {
		return "@" + c.Username
	}

	return ""
}

// InlineKeyboardMarkup | ReplyKeyboardMarkup | ReplyKeyboardRemove | ForceReply
type Markup interface{}

// InputMediaAudio | InputMediaDocument | InputMediaPhoto | InputMediaVideo
type MediaGroupInputMedia InputMedia // TODO: replace this with just []InputMedia

type NamedReader interface {
	io.Reader
	Name() string
}

// This object represents the content of a media message to be sent. It should be one of
//
// - InputMediaAnimation
//
// - InputMediaDocument
//
// - InputMediaAudio
//
// - InputMediaPhoto
//
// - InputMediaVideo
//
// https://core.telegram.org/bots/api#inputmedia
type InputMedia interface {
	setMedia(string)
	getMedia() InputFile
}
