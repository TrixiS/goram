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
	buf.Write([]byte(prefix))
	buf.Write([]byte{delim})

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
	buf := bytes.NewBufferString(callbackData)
	prefixOffset := len(prefix)
	metaLen := prefixOffset + 1 // +1 for delim
	metaBuf := make([]byte, metaLen)

	var value T

	n, _ := buf.Read(metaBuf)

	if n != metaLen || metaBuf[prefixOffset] != delim ||
		string(metaBuf[:prefixOffset]) != prefix {

		return value, ErrInvalidPrefix
	}

	decoder := base64.NewDecoder(base64.RawStdEncoding, buf)
	err := binary.Read(decoder, binary.LittleEndian, &value)
	return value, err
}

const HandlersDataKey = "callbackData"

// Creates callback query filter for callback data. Unpacks callback data and check the prefix.
// If a query has no data, the created filter returns false.
// If the prefix matches, the created filter puts unpacked callback data to handler data
// with "callbackData" key and returns true. Othersise returns false.
// Panics if the error returned from .Unpack() != ErrInvalidPrefix.
func Filter[T any](prefix string) handlers.Filter[*goram.CallbackQuery] {
	return func(ctx context.Context, bot *goram.Bot, query *goram.CallbackQuery, data handlers.Data) bool {
		if query.Data == "" {
			return false
		}

		value, err := Unpack[T](prefix, query.Data)

		if err != nil {
			if err != ErrInvalidPrefix {
				panic(err)
			}

			return false
		}

		data[HandlersDataKey] = value
		return true
	}
}

// Does the same as .Filter() but also applies the provided predicate.
func FilterFunc[T any](
	prefix string,
	predicate func(data T) bool,
) handlers.Filter[*goram.CallbackQuery] {
	return func(ctx context.Context, bot *goram.Bot, query *goram.CallbackQuery, data handlers.Data) bool {
		if query.Data == "" {
			return false
		}

		storedValue, exists := data[HandlersDataKey]

		if exists {
			if data, ok := storedValue.(T); ok {
				return predicate(data)
			}
		}

		value, err := Unpack[T](prefix, query.Data)

		if err != nil {
			if err != ErrInvalidPrefix {
				panic(err)
			}

			return false
		}

		data[HandlersDataKey] = value
		return predicate(value)
	}
}
