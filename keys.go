package mssdb

import (
	"errors"
	redigo "github.com/gomodule/redigo/redis"
)

func (rp *mPool) Get(key interface{}) (value interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return scon.Do("get", key)
}

//
func (rp *mPool) Set(key interface{}, value interface{}) (err error) {
	scon := rp.getWrite()
	defer scon.Close()
	_, err = scon.Do("set", key, value) // +OK, nil
	return
}

//
func (rp *mPool) SetEx(key interface{}, seconds int, value interface{}) (e error) {
	scon := rp.getWrite()
	defer scon.Close()
	_, e = scon.Do("setex", key, seconds, value)
	return
}

//
func (rp *mPool) SetNx(key interface{}, value interface{}) (int_value interface{}, e error) {
	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("setnx", key, value)
}

//
func (rp *mPool) Incr(key interface{}) (int64_value interface{}, err error) {
	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("incr", key)
}

func (rp *mPool) IncrBy(key interface{}, value interface{}) (int64_value interface{}, err error) {
	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("incr", key, value)
}

// 批量获取
func (rp *mPool) MGet(keys ...interface{}) (value []interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Values(scon.Do("mget", keys...))
}

/*
批量设置
< key value > 序列
*/
func (rp *mPool) MSet(kvs ...interface{}) (e error) {
	if len(kvs)%2 != 0 {
		return errors.New("mset invalid arguments number")
	}
	scon := rp.getWrite()
	defer scon.Close()
	_, e = scon.Do("mset", kvs...)
	return
}

// 返回值小于1，表示键不存在
func (rp *mPool) Del(key ...interface{}) (int_value interface{}, err error) {
	conn := rp.getWrite()
	defer conn.Close()
	return conn.Do("del", key...)
}

func (rp *mPool) Keys(pattern string) (keys []string, e error) {
	conn := rp.getRead()
	defer conn.Close()
	return redigo.Strings(conn.Do("KEYS", pattern))
}

//
func (rp *mPool) GetSet(key interface{}, value interface{}) (interface{}, error) {
	conn := rp.getWrite()
	defer conn.Close()
	return conn.Do("getset", key)
}

// todo strlen
// todo substr
// todo getrange

//
func (rp *mPool) Exists(key interface{}) (bool, error) {
	conn := rp.getRead()
	defer conn.Close()
	return redigo.Bool(conn.Do("exists", key))
}

/*
获取key的有效时间
*/
func (rp *mPool) TTL(key interface{}) (expire int64, e error) {
	conn := rp.getWrite()
	defer conn.Close()
	return redigo.Int64(conn.Do("ttl", key))
}

/*
设置key的有效时间,返回值不等于1，表示键不存在
*/
func (rp *mPool) Expire(key interface{}, expire int) (int_value interface{}, err error) {
	conn := rp.getWrite()
	defer conn.Close()
	return conn.Do("expire", key, expire)
}

//
func (rp *mPool) GetBit(key interface{}, idx int) (int, error) {
	conn := rp.getRead()
	defer conn.Close()
	return redigo.Int(conn.Do("getbit", key, idx))
}

//
func (rp *mPool) SetBit(key interface{}, idx int, b bool) (int, error) {
	conn := rp.getWrite()
	defer conn.Close()
	val := 0
	if b {
		val = 1
	}
	return redigo.Int(conn.Do("setbit", key, val))
}

//
func (rp *mPool) CountBit(key interface{}) (int, error) {
	conn := rp.getWrite()
	defer conn.Close()
	return redigo.Int(conn.Do("bitcount", key))
}

// TODO keys
