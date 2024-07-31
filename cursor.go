package gocodec

import (
	"errors"
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

func (c *Cursor[T]) Rest() ([]T, error) {
	if c.EOF() {
		return nil, errors.New("EOF")
	}
	return c.buffer[c.offset:], nil
}

func (c *Cursor[T]) Len() int64 {
	return int64(len(c.buffer) - c.offset)
}

func (c *Cursor[T]) Position() int {
	return c.offset
}

func (c *Cursor[T]) Grow(ts []T) (n int, err error) {
	c.buffer = append(c.buffer, ts...)
	return len(ts), nil
}

func (c *Cursor[T]) Undo(n int) {
	if c.offset < n {
		c.offset = 0
	} else {
		c.offset -= n
	}
}

func (c *Cursor[T]) Rewind() {
	c.offset = 0
}

func (c *Cursor[T]) Skip(n int) {
	if c.offset+n > len(c.buffer) {
		c.offset = len(c.buffer)
	} else {
		c.offset += n
	}
}

func (c *Cursor[T]) SkipTo(delim T) error {
	for i, v := range c.buffer[c.offset:] {
		if v == delim {
			c.offset += i + 1
			return nil
		}
	}
	return errors.New("no such element")
}

func (c *Cursor[T]) UnTakeN(n int) {
	if n > len(c.buffer) {
		c.offset = 0
	} else {
		c.offset -= n
		if c.offset < 0 {
			c.offset = 0
		}
	}
}

func (c *Cursor[T]) TakeN(n int) (value []T, err error) {
	return c.ReadN(n)
}

func (c *Cursor[T]) Read() (value T, err error) {
	if c.offset >= len(c.buffer) {
		err = errors.New("overflow")
		return
	}
	value = c.buffer[c.offset]
	c.offset++
	err = nil
	return
}

func (c *Cursor[T]) ReadN(n int) (value []T, err error) {
	if c.offset >= len(c.buffer) {
		return nil, errors.New("overflow")
	}

	if c.offset+n > len(c.buffer) {
		return nil, errors.New("no enough elements")
	}

	value = c.buffer[c.offset : c.offset+n]
	c.offset += n
	err = nil
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

func (c *Cursor[T]) Till(delim T) (value []T, err error) {
	if c.EOF() {
		return nil, errors.New("EOF")
	}

	for i, v := range c.buffer[c.offset:] {
		if v == delim {
			n := i + 1
			value = c.buffer[c.offset : c.offset+n]
			c.offset += n
			err = nil
			return
		}
	}
	err = errors.New("no such element")
	return
}

func (c Cursor[T]) NextN(n int) (Cursor[T], error) {
	if c.EOF() {
		return c, errors.New("EOF")
	}

	if c.offset+n > len(c.buffer) {
		return c, errors.New("no enough elements")
	}

	return Cursor[T]{
		buffer: c.buffer,
		offset: c.offset + n,
	}, nil
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
