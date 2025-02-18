package types

import (
	"io"
	"strconv"
)

// Use this struct to specify a chat id or username
type ChatId struct {
	ID       int64  // Negative int64 (-100...) for channel and some group ids, positive for user ids
	Username string // Plain username, without leading @
}

// For json encoding ChatID inside structs
func (c ChatId) MarshalJSON() ([]byte, error) {
	if c.ID != 0 {
		stringID := strconv.FormatInt(c.ID, 10)
		return []byte(stringID), nil
	}

	if c.Username != "" {
		return []byte("\"" + "@" + c.Username + "\""), nil
	}

	return nil, nil
}

func (c ChatId) String() string {
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

// Use this if you need to pass an io.Reader that does not have .Name() method.
//
// For example: you want to send a photo via Bot.SendPhoto() method, but you have only a bytes.Buffer and the photo filename.
//
// But it's always better to send a file from the file system directly without additional copying.
type NameReader struct {
	Reader   io.Reader
	FileName string
}

func (n *NameReader) Name() string {
	return n.FileName
}

func (n *NameReader) Read(b []byte) (int, error) {
	return n.Reader.Read(b)
}
