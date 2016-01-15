package shanhu

import (
	"strings"
)

func pathSplit(p string) (string, []string) {
	subs := strings.Split(p, "/")
	switch len(subs) {
	case 0:
		return "", nil
	case 1:
		return subs[0], nil
	default:
		return subs[0], subs[1:]
	}
}
