package doozer

import (
	"errors"
)

var (
	ErrNoAddrs = errors.New("no known address")
	ErrBadTag  = errors.New("bad tag")
	ErrClosed  = errors.New("closed")
)

var (
	ErrOther    Response_Err = Response_OTHER
	ErrNotDir   Response_Err = Response_NOTDIR
	ErrIsDir    Response_Err = Response_ISDIR
	ErrNoEnt    Response_Err = Response_NOENT
	ErrRange    Response_Err = Response_RANGE
	ErrOldRev   Response_Err = Response_REV_MISMATCH
	ErrTooLate  Response_Err = Response_TOO_LATE
	ErrReadonly Response_Err = Response_READONLY
)

type Error struct {
	Err    Response_Err
	Detail string
}

func newError(t *txn) *Error {
	return &Error{
		Err:    *t.resp.ErrCode,
		Detail: t.resp.GetErrDetail(),
	}
}

func (e *Error) Error() (s string) {
	s = e.Err.String()
	if e.Detail != "" {
		s += ": " + e.Detail
	}
	return s
}
