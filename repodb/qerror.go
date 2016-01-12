package repodb

import (
	"fmt"
)

type qerror struct {
	query string
	err   error
}

func (q *qerror) Error() string {
	return fmt.Sprintf("%s: query:\n%q", q.err, q.query)
}

func qerr(q string, err error) error {
	if err == nil {
		return nil
	}
	return &qerror{q, err}
}
