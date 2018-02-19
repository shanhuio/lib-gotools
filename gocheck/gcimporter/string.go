package gcimporter

func readString(p *reader) string {
	if p.debugFormat {
		readMarker(p, 's')
	}
	// if the string was seen before, i is its index (>= 0)
	// (the empty string is at index 0)
	i := readRawInt64(p)
	if i >= 0 {
		return p.strList[i]
	}
	// otherwise, i is the negative string length (< 0)
	if n := int(-i); n <= cap(p.buf) {
		p.buf = p.buf[:n]
	} else {
		p.buf = make([]byte, n)
	}
	for i := range p.buf {
		p.buf[i] = readRawByte(p)
	}
	s := string(p.buf)
	p.strList = append(p.strList, s)
	return s
}
