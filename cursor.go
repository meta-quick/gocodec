package gocodec

import (
	"fmt"
	"reflect"
)

type Cursor[T comparable] struct {
	buffer []T
	offset int
}

func NewCursor[T comparable](ts []T) Cursor[T] {
	return Cursor[T]{
		buffer: ts,
		offset: 0,
	}
}

func (c *Cursor[T]) String() string {
	return fmt.Sprintf("Cursor(%v)", c.offset)
}

func (c *Cursor[T]) EOF() bool {
	return c.offset >= len(c.buffer)
}

func (c *Cursor[T]) Rest() []T {
	if c.EOF() {
		return nil
	}
	return c.buffer[c.offset:]
}

func (c *Cursor[T]) Len() int {
	return len(c.buffer) - c.offset
}

func (c *Cursor[T]) Position() int {
	return c.offset
}

func (c *Cursor[T]) Read() (value T) {
	if c.offset >= len(c.buffer) {
		return
	}
	value = c.buffer[c.offset]
	c.offset++
	return
}

func (c *Cursor[T]) ReadN(n int) (value []T) {
	if c.offset >= len(c.buffer) {
		return
	}

	if c.offset+n > len(c.buffer) {
		n = len(c.buffer) - c.offset
	}

	value = c.buffer[c.offset : c.offset+n]
	c.offset += n
	return
}

func (c Cursor[T]) Next() Cursor[T] {
	if c.EOF() {
		return c
	}
	return Cursor[T]{
		buffer: c.buffer,
		offset: c.offset + 1,
	}
}

func (c *Cursor[T]) Till(delim T) (value []T, n int) {
	if c.EOF() {
		return nil, 0
	}

	for i, v := range c.buffer[c.offset:] {
		if v == delim {
			n = i + 1
			value = c.buffer[c.offset : c.offset+n]
			c.offset += n
			return
		}
	}
	value = c.buffer[c.offset:]
	n = len(value)
	c.offset += n
	return
}

func (c Cursor[T]) NextN(n int) Cursor[T] {
	if c.EOF() {
		return c
	}

	if c.offset+n > len(c.buffer) {
		n = len(c.buffer) - c.offset
	}

	return Cursor[T]{
		buffer: c.buffer,
		offset: c.offset + n,
	}
}

func (c *Cursor[T]) ToEOF() Cursor[T] {
	return Cursor[T]{
		buffer: c.buffer,
		offset: len(c.buffer),
	}
}

func (c *Cursor[T]) To(other Cursor[T]) []T {
	if c.EOF() || &c.buffer[0] != &other.buffer[0] {
		return nil
	}
	return c.buffer[c.offset:other.offset]
}

func (c *Cursor[T]) Addr() *T {
	if c.EOF() {
		return nil
	}
	return &c.buffer[c.offset]
}

type Span[T comparable] struct {
	Start, End Cursor[T]
}

func NewSpan[T comparable](start, end Cursor[T]) Span[T] {
	return Span[T]{
		Start: start,
		End:   end,
	}
}

func (s Span[T]) Value() []T {
	return s.Start.To(s.End)
}

func (s Span[T]) String() string {
	switch slice := reflect.ValueOf(s.Start.To(s.End)).Interface().(type) {
	case []rune:
		return fmt.Sprintf("Span(%q)", string(slice))
	default:
		return fmt.Sprintf("Span(%v...%v)", s.Start.Position(), s.End.Position())
	}
}
