package mssdb

import (
	"errors"
	redigo "github.com/gomodule/redigo/redis"
)

func (rp *SsdbPool) HGet(key interface{}, name interface{}) (value interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return scon.Do("hget", key, name)
}

// return num, error
func (rp *SsdbPool) HSet(key interface{}, id interface{}, value interface{}) (int_value interface{}, err error) {
	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("hset", key, id, value)
}

func (rp *SsdbPool) HExists(key interface{}, id interface{}) (bool, error) {
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Bool(scon.Do("hexists", key, id))
}

func (rp *SsdbPool) HLen(key interface{}) (int_value interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return scon.Do("hlen", key)
}

// int_value > 0 成功删除， = 0 id不存在
func (rp *SsdbPool) HDel(key interface{}, id interface{}) (int_value interface{}, err error) {
	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("hdel", key, id)
}

/*
HDel批量删除某个Key中的元素
	args: 第一个必须是key，后面的都是id
*/
func (rp *SsdbPool) HMDel(args ...interface{}) (int_value interface{}, err error) {
	if len(args) < 3 {
		return 0, errors.New("hmdel invalid args")
	}

	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("hmdel", args...)
}

// if need, do hdecrby
func (rp *SsdbPool) HIncrBy(key interface{}, field interface{}, increment int64) (int64_value interface{}, e error) {
	scon := rp.getWrite()
	defer scon.Close()

	return scon.Do("hincrby", key, field, increment)
}

// hkeys
func (rp *SsdbPool) HKeys(key interface{}) (value []interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()

	return redigo.Values(scon.Do("hkeys", key))
}

func (rp *SsdbPool) HVals(key interface{}) (value []interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()

	return redigo.Values(scon.Do("hvals", key))
}

/*
HMGet针对同一个key获取hashset中的部分元素的值

参数：
	args: 第一个值必须是key，后续的值都是id
	reply=>[item1, val1, item2, val2, item3, val3...]
*/
func (rp *SsdbPool) HMGet(args ...interface{}) (value []interface{}, err error) {
	if len(args) < 2 {
		return []interface{}{}, errors.New("hmget invalid args")
	}

	scon := rp.getRead()
	defer scon.Close()

	return redigo.Values(scon.Do("hmget", args...))
}

/*
HMSet针对同一个key设置hashset中的部分元素的值

参数：
	args: key item value [item2, value2...] 值对
	return: 新插入的key数量(如果item存在不加1) error
*/
func (rp *SsdbPool) HMSet(args...interface{}) (int_value interface{}, err error) {
	if len(args) < 2 {
		return nil, errors.New("hmset invalid args")
	}

	scon := rp.getWrite()
	defer scon.Close()

	return scon.Do("hmset", args...)
}

//
func (rp *SsdbPool) HGetAll(key interface{}) (value []interface{}, err error) {
	scon := rp.getRead()
	defer scon.Close()

	return redigo.Values(scon.Do("hgetall", key))
}

