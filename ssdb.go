package mssdb

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"time"
)

// NOTE!!!! when use this, must close conn, mannul
// NewRedigoPool base pool, return redigo pool
func NewRedigoPool(addr string, maxIdle, maxActive, db int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialDatabase(db))
			if err != nil {
				return nil, err
			}
			return c, err
		},

		// NOTE Comment for preference
		/*TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		*/

		Wait: true,
	}
}

type Option struct {
	DbId      int
	MaxIdle   int
	MaxActive int
	Address   string
}

//
func NewRWPool(woption Option, roptions []Option) *SsdbPool {
	rpools := make([]*redis.Pool, len(roptions))
	for idx, roption := range roptions {
		rpools[idx] = NewRedigoPool(roption.Address, roption.MaxIdle, roption.MaxActive, roption.DbId)
	}

	wpool := NewRedigoPool(woption.Address, woption.MaxIdle, woption.MaxActive, woption.DbId)

	mp := newMPool(&rw{wpool, rpools})
	go mp.check(0)

	return mp
}

// NewPool create pool, which auto close conn
func NewPool(address string, maxIdle, maxActive int) *SsdbPool {
	rp := NewRedigoPool(address, maxIdle, maxActive, 0)
	c := &common{Pool: rp}
	return newMPool(c)
}

type common struct {
	*redis.Pool
}

func (p *common) getRead() redis.Conn {
	return p.Get()
}

func (p *common) getWrite() redis.Conn {
	return p.Get()
}

func (p *common) Stat() map[string]interface{} {
	return map[string]interface{}{"master": p.ActiveCount()}
}

func (p *common) check(int) {
	// TODO
}

type rw struct {
	wPool  *redis.Pool
	rPools []*redis.Pool
}

func (p *rw) getRead() redis.Conn {
	rpool := p.rPools
	rp_size := len(rpool)
	if rp_size == 0 {
		return p.getWrite()
	}

	return rpool[rand.Int()%rp_size].Get()
}

func (p *rw) getWrite() redis.Conn {
	return p.wPool.Get()
}

//仅检查从库是否健康, 剔除掉线从库
func (p *rw) check(int) {
	rnum := len(p.rPools)
	rpool := make([]*redis.Pool, rnum)
	copy(rpool, p.rPools)

	tmp := make([]*redis.Pool, rnum)
	for {
		cnt := 0
		time.Sleep(5 * time.Second)
		for idx, rp := range rpool {
			conn := rp.Get()
			_, err := conn.Do("PING")
			if err != nil {
				fmt.Printf("Kick Slave %d [%v]\n", idx, err)
			} else {
				tmp[cnt] = rp
				cnt++
			}
			conn.Close()
		}

		if cnt > 0 { // slice alive slave
			p.rPools = tmp[:cnt]
		}
	}
}

func (rp *rw) Stat() (res map[string]interface{}) {
	res = make(map[string]interface{})
	res["master"] = rp.wPool.ActiveCount()
	for i, p := range rp.rPools {
		res[fmt.Sprintf("slave_%d", i)] = p.ActiveCount()
	}

	return
}
