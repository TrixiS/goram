package cbdata

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"errors"

	"github.com/TrixiS/goram"
	"github.com/TrixiS/goram/handlers"
)

const delim byte = ':'

var ErrInvalidPrefix = errors.New("invalid prefix")

// Packs callback data string.
//
// T should be fixed size.
// T will be binary encoded to consume less space (because callback data is up to 64 bytes only).
// Prefix is used to determine what value got encoded to the string. See .Unpack().
// Prefix should not contain ':' since it is used as a delimiter.
//
// Returns packed callback data string. Panics if T can't be binary encoded.
// The returned string is *not* checked to fit 64 bytes.
func Pack[T any](prefix string, value T) string {
	buf := &bytes.Buffer{}
	buf.WriteString(prefix)
	buf.WriteByte(delim)

	encoder := base64.NewEncoder(base64.RawStdEncoding, buf)

	if err := binary.Write(encoder, binary.LittleEndian, value); err != nil {
		panic(err)
	}

	encoder.Close()
	return buf.String()
}

// Unpacks a string encoded by .Pack().
//
// If prefixes do not match, returns ErrInvalidPrefix.
// Otherwise returns encoded value and a binary read error, if occured.
func Unpack[T any](prefix string, callbackData string) (T, error) {
	var value T

	prefixOffset := len(prefix)
	metaLen := prefixOffset + 1 // +1 for delim

	b := []byte(callbackData)

	if len(b) <= metaLen || b[prefixOffset] != delim || string(b[:prefixOffset]) != prefix {
		return value, ErrInvalidPrefix
	}

	reader := bytes.NewReader(b[metaLen:])
	decoder := base64.NewDecoder(base64.RawStdEncoding, reader)
	err := binary.Read(decoder, binary.LittleEndian, &value)
	return value, err
}

// handlers.Data key for unpacked callback data
const Key = "callbackData"

// Creates callback query filter for callback data. Unpacks callback data and check the prefix.
// If a query has no data, the created filter returns false.
// If the prefix matches, the created filter puts unpacked callback data to handler data
// with "callbackData" key and returns true. Otherwise returns false.
func Filter[T any](prefix string) handlers.Filter[*goram.CallbackQuery] {
	return func(ctx context.Context, bot *goram.Bot, query *goram.CallbackQuery, data handlers.Data) (bool, error) {
		if query.Data == "" {
			return false, nil
		}

		value, err := Unpack[T](prefix, query.Data)

		if err != nil {
			return false, err
		}

		data[Key] = value
		return true, nil
	}
}

// Does the same as .Filter() but also applies the provided predicate.
func FilterFunc[T any](
	prefix string,
	predicate func(data T) bool,
) handlers.Filter[*goram.CallbackQuery] {
	return func(ctx context.Context, bot *goram.Bot, query *goram.CallbackQuery, data handlers.Data) (bool, error) {
		if query.Data == "" {
			return false, nil
		}

		storedValue, exists := data[Key]

		if exists {
			if data, ok := storedValue.(T); ok {
				return predicate(data), nil
			}
		}

		value, err := Unpack[T](prefix, query.Data)

		if err != nil {
			return false, err
		}

		data[Key] = value
		return predicate(value), nil
	}
}
