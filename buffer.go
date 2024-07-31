package gocodec

import (
	"errors"
)

var ErrBufferFull = errors.New("funny/binary.Buffer: buffer full")

type Buffer struct {
	cursor Cursor[byte]
}

func (buf *Buffer) Error() error {
	return nil
}

func (buf *Buffer) Len() int64 {
	return buf.cursor.Len()
}

func (buf *Buffer) Reset() {
	buf.cursor.Reset()
}

func (buf *Buffer) UnTake(n int) {
	buf.cursor.UnTakeN(n)
}

func (buf *Buffer) ReadUint8() (v uint8, err error) {
	v, err = buf.cursor.Read()
	return
}

func (buf *Buffer) ReadByte() (byte, error) {
	return buf.ReadUint8()
}

func (buf *Buffer) Read(b []byte) (int, error) {
	data, err := buf.cursor.TakeN(len(b))
	if err != nil {
		return 0, err
	}
	n := copy(b, data)
	return n, nil
}

func (buf *Buffer) ReadBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	n, err := buf.Read(b)

	if err != nil {
		return nil, err
	}
	return b, nil
}

func (buf *Buffer) ReadUvarint() (uint64, error) {
	data, err := buf.cursor.TakeN(8)
	if err != nil {
		return 0, err
	}
	v, _ := GetUvarint(data)
	return v, nil
}

func (buf *Buffer) ReadVarint() (int64, error) {
	data, err := buf.cursor.TakeN(8)
	if err != nil {
		return 0, err
	}
	v, _ := GetVarint(data)
	return v, nil
}

func (buf *Buffer) ReadUint16BE() (v uint16, err error) {
	data, err := buf.cursor.TakeN(2)
	if err != nil {
		return 0, err
	}
	v = GetUint16BE(data)
	return
}

func (buf *Buffer) ReadUint16LE() (v uint16, err error) {
	data, err := buf.cursor.TakeN(2)
	if err != nil {
		return 0, err
	}

	v = GetUint16LE(data)
	return
}

func (buf *Buffer) ReadBytesTill(delim byte) (data []byte, err error) {
	data, err = buf.cursor.Till(delim)
	return
}

func (buf *Buffer) ReadString(n int) (string, error) {
	data, err := buf.ReadBytes(n)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (buf *Buffer) ReadLine() (line string, err error) {
	data, err := buf.ReadBytesTill('\n')

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (buf *Buffer) ReadUint24BE() (v uint32, err error) {
	data, err := buf.cursor.TakeN(3)
	if err != nil {
		return 0, err
	}
	v = GetUint24BE(data)
	return
}

func (buf *Buffer) ReadUint24LE() (v uint32, err error) {
	data, err := buf.cursor.TakeN(3)
	if err != nil {
		return 0, err
	}
	v = GetUint24LE(data)
	return
}

func (buf *Buffer) ReadUint32BE() (v uint32, err error) {
	data, err := buf.cursor.TakeN(4)
	if err != nil {
		return 0, err
	}
	v = GetUint32BE(data)
	return
}

func (buf *Buffer) ReadUint32LE() (v uint32, err error) {
	data, err := buf.cursor.TakeN(4)
	if err != nil {
		return 0, err
	}
	v = GetUint32LE(data)
	return
}

func (buf *Buffer) ReadUint40BE() (v uint64, err error) {
	data, err := buf.cursor.TakeN(5)
	if err != nil {
		return 0, err
	}
	v = GetUint40BE(data)
	return
}

func (buf *Buffer) ReadUint40LE() (v uint64, err error) {
	data, err := buf.cursor.TakeN(5)
	if err != nil {
		return 0, err
	}
	v = GetUint40LE(data)
	return
}

func (buf *Buffer) ReadUint48BE() (v uint64, err error) {
	data, err := buf.cursor.TakeN(6)
	if err != nil {
		return 0, err
	}
	v = GetUint48BE(data)
	return
}

func (buf *Buffer) ReadUint48LE() (v uint64, err error) {
	data, err := buf.cursor.TakeN(6)
	if err != nil {
		return 0, err
	}
	v = GetUint48LE(data)
	return
}

func (buf *Buffer) ReadUint56BE() (v uint64, err error) {
	data, err := buf.cursor.TakeN(7)
	if err != nil {
		return 0, err
	}
	v = GetUint56BE(data)
	return
}

func (buf *Buffer) ReadUint56LE() (v uint64, err error) {
	data, err := buf.cursor.TakeN(7)
	if err != nil {
		return 0, err
	}
	v = GetUint56LE(data)
	return
}

func (buf *Buffer) ReadUint64BE() (v uint64, err error) {
	data, err := buf.cursor.TakeN(8)
	if err != nil {
		return 0, err
	}
	v = GetUint64BE(data)
	return
}

func (buf *Buffer) ReadUint64LE() (v uint64, err error) {
	data, err := buf.cursor.TakeN(8)
	if err != nil {
		return 0, err
	}
	v = GetUint64LE(data)
	return
}

func (buf *Buffer) ReadFloat32BE() (v float32, err error) {
	data, err := buf.cursor.TakeN(4)
	if err != nil {
		return 0, err
	}
	v = GetFloat32BE(data)
	return
}

func (buf *Buffer) ReadFloat32LE() (v float32, err error) {
	data, err := buf.cursor.TakeN(4)
	if err != nil {
		return 0, err
	}
	v = GetFloat32LE(data)
	return
}

func (buf *Buffer) ReadFloat64BE() (v float64, err error) {
	data, err := buf.cursor.TakeN(8)
	if err != nil {
		return 0, err
	}
	v = GetFloat64BE(data)
	return
}

func (buf *Buffer) ReadFloat64LE() (v float64, err error) {
	data, err := buf.cursor.TakeN(8)
	if err != nil {
		return 0, err
	}
	v = GetFloat64LE(data)
	return
}

func (buf *Buffer) ReadInt8() (int8, error) {
	data, err := buf.ReadUint8()
	return int8(data), err
}

func (buf *Buffer) ReadInt16BE() (int16, error) {
	data, err := buf.ReadUint16BE()
	return int16(data), err
}

func (buf *Buffer) ReadInt16LE() (int16, error) {
	data, err := buf.ReadUint16LE()
	return int16(data), err
}

func (buf *Buffer) ReadInt24BE() (int32, error) {
	data, err := buf.ReadUint24BE()
	return int32(data), err
}
func (buf *Buffer) ReadInt24LE() (int32, error) {
	data, err := buf.ReadUint24LE()
	return int32(data), err
}
func (buf *Buffer) ReadInt32BE() (int32, error) {
	data, err := buf.ReadUint32BE()
	return int32(data), err
}
func (buf *Buffer) ReadInt32LE() (int32, error) {
	data, err := buf.ReadUint32LE()
	return int32(data), err
}
func (buf *Buffer) ReadInt40BE() (int64, error) {
	data, err := buf.ReadUint40BE()
	return int64(data), err
}
func (buf *Buffer) ReadInt40LE() (int64, error) {
	data, err := buf.ReadUint40LE()
	return int64(data), err
}
func (buf *Buffer) ReadInt48BE() (int64, error) {
	data, err := buf.ReadUint48BE()
	return int64(data), err
}
func (buf *Buffer) ReadInt48LE() (int64, error) {
	data, err := buf.ReadUint48LE()
	return int64(data), err
}
func (buf *Buffer) ReadInt56BE() (int64, error) {
	data, err := buf.ReadUint56BE()
	return int64(data), err
}
func (buf *Buffer) ReadInt56LE() (int64, error) {
	data, err := buf.ReadUint56LE()
	return int64(data), err
}
func (buf *Buffer) ReadInt64BE() (int64, error) {
	data, err := buf.ReadUint64BE()
	return int64(data), err
}
func (buf *Buffer) ReadInt64LE() (int64, error) {
	data, err := buf.ReadUint64LE()
	return int64(data), err
}
func (buf *Buffer) ReadIntBE() (int, error) {
	data, err := buf.ReadUint32BE()
	return int(data), err
}
func (buf *Buffer) ReadIntLE() (int, error) {
	data, err := buf.ReadUint32LE()
	return int(data), err
}
func (buf *Buffer) ReadUintBE() (uint, error) {
	data, err := buf.ReadUint32BE()
	return uint(data), err
}
func (buf *Buffer) ReadUintLE() (uint, error) {
	data, err := buf.ReadUint32LE()
	return uint(data), err
}

func (buf *Buffer) Take(n int) (data []byte, err error) {
	return buf.cursor.TakeN(n)
}

func (buf *Buffer) Write(b []byte) (int, error) {
	return buf.cursor.Grow(b)
}

func (buf *Buffer) WriteString(s string) (int, error) {
	return buf.cursor.Grow([]byte(s))
}

func (buf *Buffer) WriteLine(lines string) (int, error) {
	return buf.WriteString(lines + "\n")
}

func (buf *Buffer) WriteUvarint(v uint64) (int, error) {
	dbuf := make([]byte, 8)
	PutUvarint(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteVarint(v int64) (int, error) {
	dbuf := make([]byte, 8)
	PutVarint(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint8(v uint8) (int, error) {
	return buf.Write([]byte{v})
}

func (buf *Buffer) WriteUint16BE(v uint16) (int, error) {
	dbuf := make([]byte, 2)
	PutUint16BE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint16LE(v uint16) (int, error) {
	dbuf := make([]byte, 2)
	PutUint16LE(dbuf, v)
	return buf.Write(dbuf)
}
func (buf *Buffer) WriteUint24BE(v uint32) (int, error) {
	dbuf := make([]byte, 3)
	PutUint24BE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint24LE(v uint32) (int, error) {
	dbuf := make([]byte, 3)
	PutUint24LE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint32BE(v uint32) (int, error) {
	dbuf := make([]byte, 4)
	PutUint32BE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint32LE(v uint32) (int, error) {
	dbuf := make([]byte, 4)
	PutUint32LE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint40BE(v uint64) (int, error) {
	dbuf := make([]byte, 5)
	PutUint40BE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint40LE(v uint64) (int, error) {
	dbuf := make([]byte, 5)
	PutUint40LE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint48BE(v uint64) (int, error) {
	dbuf := make([]byte, 6)
	PutUint48BE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint48LE(v uint64) (int, error) {
	dbuf := make([]byte, 6)
	PutUint48LE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint56BE(v uint64) (int, error) {
	dbuf := make([]byte, 7)
	PutUint56BE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint56LE(v uint64) (int, error) {
	dbuf := make([]byte, 7)
	PutUint56LE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint64BE(v uint64) (int, error) {
	dbuf := make([]byte, 8)
	PutUint64BE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteUint64LE(v uint64) (int, error) {
	dbuf := make([]byte, 8)
	PutUint64LE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteFloat32BE(v float32) (int, error) {
	dbuf := make([]byte, 4)
	PutFloat32BE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteFloat32LE(v float32) (int, error) {
	dbuf := make([]byte, 4)
	PutFloat32LE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteFloat64BE(v float64) (int, error) {
	dbuf := make([]byte, 8)
	PutFloat64BE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteFloat64LE(v float64) (int, error) {
	dbuf := make([]byte, 8)
	PutFloat64LE(dbuf, v)
	return buf.Write(dbuf)
}

func (buf *Buffer) WriteInt8(v int8) (int, error) {
	return buf.WriteUint8(uint8(v))
}
func (buf *Buffer) WriteInt16BE(v int16) (int, error) { return buf.WriteUint16BE(uint16(v)) }
func (buf *Buffer) WriteInt16LE(v int16) (int, error) { return buf.WriteUint16LE(uint16(v)) }
func (buf *Buffer) WriteInt24BE(v int32) (int, error) { return buf.WriteUint24BE(uint32(v)) }
func (buf *Buffer) WriteInt24LE(v int32) (int, error) { return buf.WriteUint24LE(uint32(v)) }
func (buf *Buffer) WriteInt32BE(v int32) (int, error) { return buf.WriteUint32BE(uint32(v)) }
func (buf *Buffer) WriteInt32LE(v int32) (int, error) { return buf.WriteUint32LE(uint32(v)) }
func (buf *Buffer) WriteInt40BE(v int64) (int, error) { return buf.WriteUint40BE(uint64(v)) }
func (buf *Buffer) WriteInt40LE(v int64) (int, error) { return buf.WriteUint40LE(uint64(v)) }
func (buf *Buffer) WriteInt48BE(v int64) (int, error) { return buf.WriteUint48BE(uint64(v)) }
func (buf *Buffer) WriteInt48LE(v int64) (int, error) { return buf.WriteUint48LE(uint64(v)) }
func (buf *Buffer) WriteInt56BE(v int64) (int, error) { return buf.WriteUint56BE(uint64(v)) }
func (buf *Buffer) WriteInt56LE(v int64) (int, error) { return buf.WriteUint56LE(uint64(v)) }
func (buf *Buffer) WriteInt64BE(v int64) (int, error) { return buf.WriteUint64BE(uint64(v)) }
func (buf *Buffer) WriteInt64LE(v int64) (int, error) { return buf.WriteUint64LE(uint64(v)) }
func (buf *Buffer) WriteIntBE(v int) (int, error)     { return buf.WriteUint64BE(uint64(v)) }
func (buf *Buffer) WriteIntLE(v int) (int, error)     { return buf.WriteUint64LE(uint64(v)) }
func (buf *Buffer) WriteUintBE(v uint) (int, error)   { return buf.WriteUint64BE(uint64(v)) }
func (buf *Buffer) WriteUintLE(v uint) (int, error)   { return buf.WriteUint64LE(uint64(v)) }
