package mssdb

//import (
//	"errors"
//	redigo "github.com/gomodule/redigo/redis"
//)
//
//// args : key, val1, val2, val3....
//func (rp *mPool) SAdd(db int, args ...interface{}) (int64, error) {
//	scon := rp.getWrite(db)
//	defer scon.Close()
//	return redigo.Int64(scon.Do("SADD", args...))
//}
//
//// args : key, val1, val2, val3...
//func (rp *mPool) SRem(db int, args ...interface{}) (int64, error) {
//	scon := rp.getWrite(db)
//	defer scon.Close()
//	return redigo.Int64(scon.Do("SREM", args...))
//}
//
//func (rp *mPool) SIsMember(db int, key interface{}, value interface{}) (isMember bool, e error) {
//	scon := rp.getRead(db)
//	defer scon.Close()
//	return redigo.Bool(scon.Do("SISMEMBER", key, value))
//}
//
///*
//SMembers获取某个key下的所有元素
//
//参数：
//	values: 必须是数组的引用
//*/
//func (rp *mPool) SMembers(db int, key interface{}) (values []interface{}, e error) {
//	scon := rp.getRead(db)
//	defer scon.Close()
//	return redigo.Values(scon.Do("SMEMBERS", key))
//}
//
///*
//SCard获取某个key下的元素数量
//
//参数：
//	values: 必须是数组的引用
//*/
//func (rp *mPool) SCard(db int, key interface{}) (count int64, e error) {
//	scon := rp.getRead(db)
//	defer scon.Close()
//	return redigo.Int64(scon.Do("SCARD", key))
//}
//
///*
//SRandMembers获取某个key下的随机count 个元素
//
//参数：
//	values: 必须是数组的引用
//*/
//func (rp *mPool) SRandMembers(db int, key interface{}, count int) (values interface{}, e error) {
//	scon := rp.getRead(db)
//	defer scon.Close()
//	return redigo.Values(scon.Do("SRANDMEMBER", key, count))
//}
//
///*
//批量添加到set类型的表中
//
//	db: 数据库表ID
//	args: 必须是<key,id>的列表
//*/
//func (rp *mPool) SMultiAdd(db int, args ...interface{}) error {
//	if len(args)%2 != 0 {
//		return errors.New("invalid arguments number")
//	}
//
//	fcon := rp.getWrite(db)
//	defer fcon.Close()
//	if e := fcon.Send("MULTI"); e != nil {
//		return e
//	}
//
//	for i := 0; i < len(args); i += 2 {
//		if e := fcon.Send("SADD", args[i], args[i+1]); e != nil {
//			fcon.Send("DISCARD")
//			return e
//		}
//	}
//	if _, e := fcon.Do("EXEC"); e != nil {
//		fcon.Send("DISCARD")
//		return e
//	}
//	return nil
//}
//
///*
//批量删除set类型表中的元素
//
//	db: 数据库表ID
//	args: 必须是<key,id>的列表
//*/
//func (rp *mPool) SMultiRem(db int, args ...interface{}) error {
//	if len(args)%2 != 0 {
//		return errors.New("invalid arguments number")
//	}
//
//	fcon := rp.getWrite(db)
//	defer fcon.Close()
//	if e := fcon.Send("MULTI"); e != nil {
//		return e
//	}
//
//	for i := 0; i < len(args); i += 2 {
//		if e := fcon.Send("SREM", args[i], args[i+1]); e != nil {
//			fcon.Send("DISCARD")
//			return e
//		}
//	}
//	if _, e := fcon.Do("EXEC"); e != nil {
//		fcon.Send("DISCARD")
//		return e
//	}
//	return nil
//}
