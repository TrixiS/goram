package goram

import (
	"io"
	"strconv"
)

// Use this struct to specify a chat id or username
type ChatId struct {
	Id       int64  // Negative int64 (-100...) for channel and some group ids, positive for user ids
	Username string // Plain username, without leading @
}

// For json encoding ChatId inside structs
func (c ChatId) MarshalJSON() ([]byte, error) {
	if c.Id != 0 {
		stringId := strconv.FormatInt(c.Id, 10)
		return []byte(stringId), nil
	}

	if c.Username != "" {
		return []byte("\"@" + c.Username + "\""), nil
	}

	return nil, nil
}

func (c ChatId) String() string {
	if c.Id != 0 {
		stringID := strconv.FormatInt(c.Id, 10)
		return stringID
	}

	if c.Username != "" {
		return "@" + c.Username
	}

	return ""
}

// InlineKeyboardMarkup | ReplyKeyboardMarkup | ReplyKeyboardRemove | ForceReply
type Markup interface{}

type NamedReader interface {
	io.Reader
	Name() string
}

// This object represents the contents of a file to be uploaded.
//
// You can use file id of existing file or any struct that implements NamedReader interface.
// Also you can use a url in the file id field to send files from the internet.
// If FileId and Reader are both set, file id will be used.
//
// See goram.NameReader also.
type InputFile struct {
	FileId string
	Reader NamedReader
}

func (i InputFile) MarshalJSON() ([]byte, error) {
	return []byte(`"` + i.FileId + `"`), nil
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

// Use this if you need to pass an io.Reader that does not have .Name() method to InputFile.
//
// For example: you want to send a photo via Bot.SendPhoto() method, but you have only a bytes.Buffer and the photo filename.
// Or you have a file, but you want different filename.
type NameReader struct {
	Reader   io.Reader
	FileName string
}

func (n NameReader) Name() string {
	return n.FileName
}

func (n NameReader) Read(b []byte) (int, error) {
	return n.Reader.Read(b)
}

func (m *Message) ChatId() ChatId {
	return ChatId{Id: m.Chat.Id}
}

func (u *User) ChatId() ChatId {
	return ChatId{Id: u.Id}
}

func (c *CallbackQuery) ChatId() ChatId {
	if c.Message != nil && c.Message.Chat != nil {
		return ChatId{Id: c.Message.Chat.Id}
	}

	return ChatId{Id: c.From.Id}
}

func (c *Chat) ChatId() ChatId {
	return ChatId{Id: c.Id}
}
