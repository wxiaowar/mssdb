package mssdb

import "github.com/gomodule/redigo/redis"

type proxy interface {
	Stat() map[string]interface{}
	getRead() redis.Conn
	getWrite() redis.Conn
	check(int)
}

type SsdbPool struct {
	proxy
}

func newMPool(rw proxy) *SsdbPool {
	return &SsdbPool{rw}
}



