package keyboards

import "github.com/TrixiS/goram"

type keyboardButton interface {
	goram.KeyboardButton | goram.InlineKeyboardButton
}

// Keyboard builder for goram.ReplyKeyboardMarkup and goram.InlineKeyboardMarkup.
//
// The generic type should be goram.KeyboardButton or goram.InlineKeyboardButton.
type Builder[B keyboardButton] struct {
	rows [][]B
}

// Creates and returns a pointer to keyboards.Builder. This is useful for chaining.
//
// Specified rows will be used as initial rows of the returned builder.
// You can pass nil as a row to specify a break.
//
// See keyboards.Builder for more info.
func NewBuilder[B keyboardButton](rows ...[]B) *Builder[B] {
	return &Builder[B]{rows}
}

// Appends a button to the last keyboard row.
func (b *Builder[B]) Add(button B) *Builder[B] {
	lastIdx := len(b.rows) - 1

	if len(b.rows) == 0 || b.rows[lastIdx] == nil {
		b.rows = append(b.rows, []B{button})
	} else {
		b.rows[lastIdx] = append(b.rows[lastIdx], button)
	}

	return b
}

// Appends a row of specified buttons to the keyboard.
func (b *Builder[B]) Row(buttons ...B) *Builder[B] {
	b.rows = append(b.rows, buttons)
	return b
}

// Appends other keyboard builder rows to this builder keyboard. Other builder remains unchanged.
func (b *Builder[B]) Merge(other *Builder[B]) *Builder[B] {
	b.rows = append(b.rows, other.rows...)
	return b
}

// Appends a break to the keyboard. You must call .Adjust() before .Build() if you use .Break().
func (b *Builder[B]) Break() *Builder[B] {
	b.rows = append(b.rows, nil)
	return b
}

// Resizes each row in the keyboard to fit at most rowSize buttons.
func (b *Builder[B]) Adjust(rowSize int) *Builder[B] {
	if rowSize <= 0 || len(b.rows) == 0 {
		return b
	}

	newRows := make([][]B, 0, len(b.rows))
	currentRow := []B{}

	for i, row := range b.rows {
		shouldFlush := i < len(b.rows)-1 && b.rows[i+1] == nil

		for _, v := range row {
			if currentRow == nil {
				currentRow = make([]B, 0, rowSize)
			}

			currentRow = append(currentRow, v)

			if len(currentRow) == rowSize {
				newRows = append(newRows, currentRow)
				currentRow = nil
			}
		}

		if shouldFlush && len(currentRow) > 0 {
			newRows = append(newRows, currentRow)
			currentRow = nil
		}
	}

	if len(currentRow) > 0 {
		newRows = append(newRows, currentRow)
	}

	b.rows = newRows
	return b
}

// Returns the built keyboard. Sets the underlying keyboard to nil so that the builder could be reused.
func (b *Builder[B]) Build() [][]B {
	rows := b.rows
	b.rows = nil
	return rows
}
