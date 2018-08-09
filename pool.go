package mssdb

import "github.com/gomodule/redigo/redis"

type proxy interface {
	Stat() map[string]interface{}
	getRead() redis.Conn
	getWrite() redis.Conn
	check(int)
}

type mPool struct {
	proxy
}

func newMPool(rw proxy) *mPool {
	return &mPool{rw}
}



