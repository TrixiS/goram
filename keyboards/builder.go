package keyboards

// Keyboard builder for goram.ReplyKeyboardMarkup and goram.InlineKeyboardMarkup.
//
// The generic type should be goram.KeyboardButton or goram.InlineKeyboardButton.
type Builder[B any] struct {
	rows [][]B
}

// Creates and returns a pointer to keyboards.Builder. This is useful for chaining.
//
// Specified rows will be used as initial rows of the returned builder.
//
// See keyboards.Builder for more info.
func NewBuilder[B any](rows ...[]B) *Builder[B] {
	return &Builder[B]{rows}
}

// Appends a button to the last keyboard row.
func (b *Builder[B]) Add(button B) *Builder[B] {
	if len(b.rows) == 0 {
		b.rows = append(b.rows, []B{button})
	} else {
		lastIdx := len(b.rows) - 1
		b.rows[lastIdx] = append(b.rows[lastIdx], button)
	}

	return b
}

// Appends a row of specified buttons to the keyboard.
func (b *Builder[B]) Row(buttons ...B) *Builder[B] {
	b.rows = append(b.rows, buttons)
	return b
}

// Resizes each row in the keyboard to fit at most rowSize buttons.
func (b *Builder[B]) Adjust(rowSize int) *Builder[B] {
	newRows := make([][]B, 0, len(b.rows))
	currentRow := make([]B, 0, rowSize)

	lastIdx := len(b.rows) - 1

rowsLoop:
	for i, row := range b.rows {
		added := 0

		for added < len(row) {
			toAdd := min(rowSize-len(currentRow), len(row)-added)

			currentRow = append(currentRow, row[added:added+toAdd]...)

			added += toAdd
			isLastAdd := i == lastIdx && added == len(row)

			if len(currentRow) == rowSize || isLastAdd {
				newRows = append(newRows, currentRow)

				if isLastAdd {
					break rowsLoop
				}

				currentRow = make([]B, 0, rowSize)
			}
		}
	}

	b.rows = newRows
	return b
}

// Returns built keyboard. Sets the underlying keyboard to nil so that the builder could be reused.
func (b *Builder[B]) Build() [][]B {
	rows := b.rows
	b.rows = nil
	return rows
}
