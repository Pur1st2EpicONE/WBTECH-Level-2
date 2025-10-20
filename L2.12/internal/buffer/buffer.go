// Package buffer implements a simple circular queue for storing lines
// and their numbers, used for context in grep operations.
package buffer

// Buffer stores lines, their numbers, and printed status in a circular queue.
type Buffer struct {
	Buffer  []string // Stored lines
	Printed []bool   // Flags indicating if lines have been printed
	Numbers []int    // Line numbers corresponding to Buffer
	Head    int      // Index of the oldest element
	Tail    int      // Index where the next element will be inserted
	Size    int      // Number of elements in the queue
}

// NewBuffer creates a new circular queue of the given size.
func NewBuffer(size int) *Buffer {
	return &Buffer{
		Buffer:  make([]string, size),
		Printed: make([]bool, size),
		Numbers: make([]int, size),
	}
}

// Enqueue adds a line and its number to the circular queue, overwriting old elements if full.
// Returns the index where the line was inserted.
func (b *Buffer) Enqueue(line string, num int) int {
	idx := b.Tail
	if b.Size == len(b.Buffer) {
		b.Buffer[idx] = line
		b.Numbers[idx] = num
		b.Printed[idx] = false
		b.Tail = b.moveIndex(b.Tail)
		b.Head = b.moveIndex(b.Head)
	} else {
		b.Buffer[idx] = line
		b.Numbers[idx] = num
		b.Printed[idx] = false
		b.Tail = b.moveIndex(b.Tail)
		b.Size++
	}
	return idx
}

// moveIndex advances an index in the circular queue, wrapping around if necessary.
func (b *Buffer) moveIndex(i int) int {
	i++
	if i == len(b.Buffer) {
		i = 0
	}
	return i
}
