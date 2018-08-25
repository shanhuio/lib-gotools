package gcimporter

func readInt64(p *reader) int64 {
	if p.debugFormat {
		readMarker(p, 'i')
	}

	return readRawInt64(p)
}

func readInt(p *reader) int {
	x := readInt64(p)
	if int64(int(x)) != x {
		errorf("exported integer too large")
	}
	return int(x)
}
