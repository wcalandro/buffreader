// Package buffreader provides a helper to read bits in a file, helpful for parsing
package buffreader

import (
	"bytes"
	"encoding/binary"
)

type Buffer struct {
	Bytes  []byte
	Offset int64 // In bytes
}

func NewBuffer(bytes []byte) Buffer {
	return Buffer{
		Bytes:  bytes,
		Offset: 0,
	}
}

func (buf Buffer) ReadInt32() (ret int32) {
	binary.Read(bytes.NewBuffer(buf.Bytes[buf.Offset:buf.Offset+4]), binary.LittleEndian, &ret)
	// Add 32 bits to offset
	buf.Offset = buf.Offset + 4
	return
}
