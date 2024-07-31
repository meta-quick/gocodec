package gocodec

import "testing"

func Test_cursor(t *testing.T) {

	bytes := []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf, '\n'}

	cursor := NewCursor(bytes)
	b, _ := cursor.Till('\n')

	println(b)
}

func Test_cursor01(t *testing.T) {

	bytes := string("hello \nworld")

	cursor := NewCursor([]byte(bytes))
	b, _ := cursor.Till('\n')

	println(string(b))
}
