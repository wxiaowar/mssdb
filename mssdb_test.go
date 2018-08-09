package mssdb

import (
	"testing"
	"fmt"
)

var p *mPool

func init() {
	p = NewPool("127.0.0.1:8888", 2, 2, 0)
}

//
func TestHMgetSet(t *testing.T) {
	return
	mk := "hkey"
	keys := []string{mk}
	for i:=0; i<5; i++ {
		key := fmt.Sprintf("key_%d", i)
		keys = append(keys, key)
		//val := fmt.Sprintf("val_%d", i)
		p.HSet(mk, key, i)
	}

	setReply, err := p.HMSet(mk, "key_1_1", 21, "key_1_2", 19)
	fmt.Printf("hmset %v %v\n", setReply, err)


	reply, err := p.HMGet(mk, "key_0", "key_1", "key_1_1", "key_1_2")
	for idx := range reply {
		bts := reply[idx].([]byte)
		fmt.Printf("idx[%d] %s = %v\n", idx, string(bts), reply[idx])
	}

	reply, err = p.HGetAll(mk)
	for idx := range reply {
		bts := reply[idx].([]byte)
		fmt.Printf("idx[%d] %s = %v\n", idx, string(bts), reply[idx])
	}

	ln, err := p.HLen(mk)
	fmt.Printf("hlen %v %v\n", ln, err)

	b, err := p.HExists(mk, "key_1")
	fmt.Printf("hexists2 %v %v\n", b, err)

	hkeys, err := p.HKeys(mk)
	fmt.Printf("hkeys %v %v\n", hkeys, err)

	hvals, err := p.HVals(mk)
	fmt.Printf("hvals %v %v\n", hvals, err)

	delReply, err := p.HMDel(mk, "key_2", "key_3")
	fmt.Printf("hmdel %v %v\n", delReply, err)

	delReply, err = p.HDel(mk, "key_1")
	fmt.Printf("hdel %v %v\n", delReply, err)

	b, err = p.HExists(mk, "key_1")
	fmt.Printf("hexists2 %v %v", b, err)
}


func TestZset(t *testing.T) {
	zkey := "zkey"
	zkey1 := "zkey1"
	zkey2 := "zkey2"
	reply, err := p.ZAdd(zkey, 100, zkey1)
	p.ZAdd(zkey, 180, "zk2")
	p.ZAdd(zkey, 200, "zk3")
	fmt.Printf("zadd reply=%v err=%v\n", reply, err)

	c, err := p.ZCount(zkey, 100, 200)
	fmt.Printf("zcount %v-%v\n", c, err)

	zrem, err := p.ZRem(zkey, "zk3")
	fmt.Printf("zrem %v-%v\n", zrem, err)

	reply, err = p.ZAdd(zkey, 110, zkey1)
	fmt.Printf("zaddopt reply=%v err=%v\n", reply, err)

	reply, err = p.ZAdd(zkey, 180, zkey2)
	fmt.Printf("zaddopt2 reply=%v err=%v\n", reply, err)

	nscore, err := p.ZIncrBy(zkey, 30, zkey1)
	fmt.Printf("zincrby %v-%v\n", nscore, err)

	score, err := p.ZScore(zkey, zkey1)
	fmt.Printf("zscore %v-%v\n", score, err)

	num, err := p.ZCard(zkey)
	fmt.Printf("zcard %v-%v\n", num, err)

	rank, err := p.ZRank(zkey, zkey1)
	fmt.Printf("zrank %v-%v\n", rank, err)

	rank, err = p.ZRevRank(zkey, zkey1)
	fmt.Printf("zrevrank %v-%v\n", rank, err)

	zrange, err := p.ZRange(zkey, 0, 3)
	fmt.Printf("zrange %v-%v\n", zrange, err)

	zrevrange, err := p.ZRevRange(zkey, 0, 3)
	fmt.Printf("zrevrange %v-%v\n", zrevrange, err)

	zranges, err := p.ZRangeWithScore(zkey, 0, 3)
	fmt.Printf("zrangeScore( %v-%v\n", zranges, err)

	zrevranges, err := p.ZRevRangeWithScore(zkey, 0, 3)
	fmt.Printf("zrevrangeScore( %v-%v\n", zrevranges, err)


	//reply, err = p.HIncrBy(hkey, key2, 10)
	//r, err := redis.Int(reply, err)
	//fmt.Printf("HIncr reply=%v err=%v\n", reply, err)
}

//
func TestRwPool(t *testing.T) {
	//wop := Option{
	//	DbId: 0,
	//	Address:"127.0.0.1:8888",
	//	MaxIdle: 2,
	//	MaxActive: 2,
	//}
	//
	//rop := []Option{
	//	{
	//		DbId: 0,
	//		Address:"127.0.0.1:8881",
	//		MaxIdle: 2,
	//		MaxActive: 2,
	//	},
	//}
	//
	//p := NewRWPool(wop, rop)
	//
	//mk := "hkey2"
	//for i:=10; i<15; i++ {
	//	key := fmt.Sprintf("key_%d", i)
	//	//val := fmt.Sprintf("val_%d", i)
	//	p.HSet(0, mk, key, i)
	//	fmt.Printf("hset %s->%d\n", key, i)
	//}
	//
	//
	//reply, err := p.HGetAll(0, mk)
	//if err != nil {
	//	fmt.Printf("hgetall %v\n", err)
	//	return
	//}
	//
	//fmt.Sprintf("reply = %v\n", reply)
	//for idx := range reply {
	//	bts := reply[idx].([]byte)
	//	fmt.Printf("idx[%d] %s = %v\n", idx, string(bts), reply[idx])
	//}
}
