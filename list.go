package mssdb

import (
	"errors"
	redigo "github.com/gomodule/redigo/redis"
)

//批量插入队尾
// args : key <val1, val2, val3...>
func (rp *mPool) RPush(args ...interface{}) (length interface{}, err error) {
	if len(args) < 2 {
		return 0, errors.New("rpush invalid args")
	}
	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("rpush", args...)
}

// 队头弹出队列数据
func (rp *mPool) LPop(key interface{}) (value interface{}, e error) {
	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("lpop", key)
}

//批量插入队头
// args : key <val1, val2, val3...>
func (rp *mPool) LPush(args ...interface{}) (length interface{}, e error) {
	if len(args) < 2 {
		return 0, errors.New("lpush invald args")
	}

	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("lpush", args...)
}

//从对尾pop
func (rp *mPool) RPop(key interface{}) (value interface{}, e error) {
	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("rpop", key)
}

/*
获取队列数据
*/
func (rp *mPool) LRange(key interface{}, start, stop interface{}) (value []interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Values(scon.Do("lrange", key, start, stop))
}

/*
获取队列长度，如果key不存在，length=0，不会报错。
*/
func (rp *mPool) LLen(key interface{}) (length int64, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Int64(scon.Do("llen", key))
}

/*
索引元素
*/
func (rp *mPool) LIndex(key interface{}, index int) (value interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return scon.Do("lindex", key, index)
}

/*
更新
*/
func (rp *mPool) LSet(key interface{}, idx int, data interface{}) (e error) {
	scon := rp.getWrite()
	defer scon.Close()
	_, e = scon.Do("lset", key, idx, data)
	return
}

/*
获取头部元素
*/
func (rp *mPool) LFront(key interface{}) (value interface{}, e error) {
	return rp.LIndex(key, 0)
}

/*
获取尾部元素
*/
func (rp *mPool) LBack(key interface{}) (value interface{}, e error) {
	return rp.LIndex(key, -1)
}
