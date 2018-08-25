package gcimporter

func readMarker(p *reader, want byte) {
	if got := readRawByte(p); got != want {
		errorf(
			"incorrect marker: got %c; want %c (pos = %d)",
			got, want, p.read,
		)
	}

	pos := p.read
	if n := int(readRawInt64(p)); n != pos {
		errorf("incorrect position: got %d; want %d", n, pos)
	}
}
