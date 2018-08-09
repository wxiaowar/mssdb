package mssdb

// not use for convert

import (
	redigo "github.com/gomodule/redigo/redis"
)

type Reply struct {
	D   interface{}
	Err error
}

func NewReply(reply interface{}, err error) *Reply {
	return &Reply{
		D:   reply,
		Err: err,
	}
}

func (r *Reply) IsError() bool {
	return r.Err != nil
}

func (r *Reply) Error() error {
	return r.Err
}

func (r *Reply) Data() interface{} {
	return r.D
}

func (r *Reply) ToInt() (int, error) {
	return redigo.Int(r.D, r.Err)
}

func (r *Reply) ToInt64() (int64, error) {
	return redigo.Int64(r.D, r.Err)
}

func (r *Reply) ToUint64() (uint64, error) {
	return redigo.Uint64(r.D, r.Err)
}

func (r *Reply) ToString() (string, error) {
	return redigo.String(r.D, r.Err)
}

func (r *Reply) ToBytes() ([]byte, error) {
	return redigo.Bytes(r.D, r.Err)
}
