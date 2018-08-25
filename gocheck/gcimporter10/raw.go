package gcimporter

import (
	"encoding/binary"
)

// readRawStringln should only be used to read the initial version string.
func readRawStringln(p *reader, b byte) string {
	p.buf = p.buf[:0]
	for b != '\n' {
		p.buf = append(p.buf, b)
		b = readRawByte(p)
	}
	return string(p.buf)
}

func readRawInt64(p *reader) int64 {
	i, err := binary.ReadVarint(byteReader{p})
	if err != nil {
		errorf("read error: %v", err)
	}
	return i
}

// byte is the bottleneck interface for reading p.data.
// It unescapes '|' 'S' to '$' and '|' '|' to '|'.
// readRawByte should only be used by low-level decoders.
func readRawByte(p *reader) byte {
	b := p.data[0]
	r := 1
	if b == '|' {
		b = p.data[1]
		r = 2
		switch b {
		case 'S':
			b = '$'
		case '|':
			// nothing to do
		default:
			errorf("unexpected escape sequence in export data")
		}
	}
	p.data = p.data[r:]
	p.read += r
	return b
}

type byteReader struct {
	*reader
}

func (r byteReader) ReadByte() (byte, error) {
	return readRawByte(r.reader), nil
}
