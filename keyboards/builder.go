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
func (k *Builder[B]) Add(button B) *Builder[B] {
	if len(k.rows) == 0 {
		k.rows = append(k.rows, []B{button})
	} else {
		lastIdx := len(k.rows) - 1
		k.rows[lastIdx] = append(k.rows[lastIdx], button)
	}

	return k
}

// Appends a row of specified buttons to the keyboard.
func (k *Builder[B]) Row(buttons ...B) *Builder[B] {
	k.rows = append(k.rows, buttons)
	return k
}

// Resizes each row in the keyboard to fit at most rowSize buttons.
func (k *Builder[B]) Adjust(rowSize int) *Builder[B] {
	newRows := make([][]B, 0, len(k.rows))
	currentRow := make([]B, 0, rowSize)

	lastIdx := len(k.rows) - 1

rowsLoop:
	for i, row := range k.rows {
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

	k.rows = newRows
	return k
}

// Returns built keyboard. Sets the underlying keyboard to nil so that the builder could be reused.
func (k *Builder[B]) Build() [][]B {
	rows := k.rows
	k.rows = nil
	return rows
}
