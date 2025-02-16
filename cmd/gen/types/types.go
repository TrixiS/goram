package types

import (
	"strconv"
)

type ChatID struct {
	ID       int64
	Username string
}

func (c ChatID) MarshalJSON() ([]byte, error) {
	if c.ID != 0 {
		stringID := strconv.FormatInt(c.ID, 10)
		return []byte(stringID), nil
	}

	if c.Username != "" {
		return []byte(c.Username), nil
	}

	panic("invalid id")
}
