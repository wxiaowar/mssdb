package mssdb

import (
	"errors"
	redigo "github.com/gomodule/redigo/redis"
	"math"
)

/*
批量添加到sorted set类型的表中
	db: 数据库表ID
	args: 必须是key <score,id>...的列表
	return： 数量，error
*/
func (rp *SsdbPool) ZAdd(args ...interface{}) (interface{}, error) {
	if (len(args)-1)%2 != 0 {
		return 0, errors.New("zadd invalid arguments number")
	}
	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("zadd", args...)
}

///*
//批量添加到sorted set类型的表中
//	opt: 可选参数，必须是NX|XX|CH|INCR|""中的一个
//	args:顺序 key opt <score,id>...列表
//	return： 数量，error
//*/
//func (rp *mPool) ZAddOpt(args ...interface{}) (interface{}, error) {
//	if (len(args)-2)%2 != 0 {
//		return 0, errors.New("zaddopt invalid arguments number")
//	}
//
//	scon := rp.getWrite()
//	defer scon.Close()
//	return scon.Do("zadd", args...)
//}

func (rp *SsdbPool) ZScore(key interface{}, id interface{}) (score int64, e error) {
	conn := rp.getRead()
	defer conn.Close()
	return redigo.Int64(conn.Do("zscore", key, id))
}

//ZIsMember判断是否是有序集合的成员
func (rp *SsdbPool) ZIsMember(key interface{}, id interface{}) (isMember bool, e error) {
	conn := rp.getRead()
	defer conn.Close()
	_, e = redigo.Float64(conn.Do("zscore", key, id))
	switch e {
	case nil:
		return true, nil
	case redigo.ErrNil:
		return false, nil
	default:
		return false, e
	}
}

//ZRem批量删除sorted set表中的元素
//
//参数：
//	db: 数据库表ID
//	args: key <id>的列表
//返回值：
//	: 每条命令影响的行数
func (rp *SsdbPool) ZRem(args ...interface{}) (interface{}, error) {
	if len(args) < 2 {
		return 0, errors.New("zrem invalid arguments number")
	}

	scon := rp.getWrite()
	defer scon.Close()
	return scon.Do("zrem", args...)
}

// key, incr, id
func (rp *SsdbPool) ZIncrBy(key interface{}, increment interface{}, id interface{}) (int64_score interface{}, e error) {
	conn := rp.getWrite()
	defer conn.Close()
	return conn.Do("zincrby", key, increment, id)
}

// return number between min max
func (rp *SsdbPool) ZCount(key interface{}, min, max float64) (count int64, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Int64(scon.Do("zcount", key, min, max))
}

// TODO zsum/zavg

// return total number
func (rp *SsdbPool) ZCard(key interface{}) (num int64, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Int64(scon.Do("zcard", key))
}

//升序(从0开始)
func (rp *SsdbPool) ZRank(key interface{}, id interface{}) (rank int64, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Int64(scon.Do("zrank", key, id))
}

//降序
func (rp *SsdbPool) ZRevRank(key interface{}, id interface{}) (rank int64, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Int64(scon.Do("zrevrank", key, id))
}

//升序 [start_rank, end_rank)
func (rp *SsdbPool) ZRange(key interface{}, start, end int) (reply []interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Values(scon.Do("zrange", key, start, end))
}

//降序 [start_rank, end_rank)
func (rp *SsdbPool) ZRevRange(key interface{}, start, end int) (reply []interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Values(scon.Do("zrevrange", key, start, end))
}

////分页获取
//func (rp *mPool) ZRangePS(key interface{}, cur int, ps int) (reply interface{}, e error) {
//	start, end := buildRange(cur, ps)
//	return rp.ZRange(key, start, end)
//}
//
////分页获取
//func (rp *mPool) ZRevRangePS(key interface{}, cur int, ps int) (reply interface{}, e error) {
//	start, end := buildRange(cur, ps)
//	return rp.ZRevRange(key, start, end)
//}

//获取SortedSet的ID集合升序
func (rp *SsdbPool) ZRangeWithScore(key interface{}, start, end int) (reply []interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()

	//if e = redigo.ScanSlice(values, &items); e != nil {
	return redigo.Values(scon.Do("zrange", key, start, end, "withscores"))
}

//获取SortedSet的ID集合降序
func (rp *SsdbPool) ZRevRangeWithScore(key interface{}, start, end int) (reply []interface{}, e error) {
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Values(scon.Do("zrevrange", key, start, end, "withscores"))
}

//func (rp *mPool) ZRangeWithScoresPS(key interface{}, cur int, ps int) (reply []interface{}, e error) {
//	if ps > 100 {
//		ps = 100
//	}
//	start, end := buildRange(cur, ps)
//	scon := rp.getRead()
//	defer scon.Close()
//	return rp.ZRangeWithScore(key, start, end)
//}
//
//func (rp *mPool) ZRevRangeWithScoresPS(key interface{}, cur int, ps int) (reply []interface{}, e error) {
//	if ps > 100 {
//		ps = 100
//	}
//	start, end := buildRange(cur, ps)
//	scon := rp.getRead()
//	defer scon.Close()
//	return rp.ZRevRangeWithScore(key, start, end)
//}

//
func (rp *SsdbPool) ZRangeByScore(key interface{}, min, max int64, limit int) (reply []interface{}, e error) {
	if limit <= 0 {
		limit = math.MaxInt32
	}
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Values(scon.Do("zrangebyscore", key, min, max, "limit", 0, limit))
}

func (rp *SsdbPool) ZRevRangeByScore(key interface{}, min, max int64, limit int) (reply []interface{}, e error) {
	if limit <= 0 {
		limit = math.MaxInt32
	}
	scon := rp.getRead()
	defer scon.Close()
	return redigo.Values(scon.Do("zrevrangebyscore", key, max, min, "limit", 0, limit))
}

// min <=score <= max 升序
func (rp *SsdbPool) ZRangeByScoreWithScore(key interface{}, min, max int64, limit int) (reply []interface{}, e error) {
	if limit <= 0 {
		limit = math.MaxInt32
	}
	scon := rp.getRead()
	defer scon.Close()
	//		s1 = fmt.Sprintf("(%d", min)
	//		cmd = "ZRANGEBYSCORE"
	return redigo.Values(scon.Do("zrangebyscore", key, min, max, "withscores", "limit", 0, limit))
}

// min <= score <= max 降序
func (rp *SsdbPool) ZRevRangeByScoreWithScore(key interface{}, min, max int64, limit int) (reply []interface{}, e error) {
	if limit <= 0 {
		limit = math.MaxInt32
	}

	scon := rp.getRead()
	defer scon.Close()
	//	s1 = fmt.Sprintf("%d", max)
	//	cmd = "ZREVRANGEBYSCORE"
	//	return redigo.Values(scon.Do(cmd, key, s1, max, "WITHSCORES", "LIMIT", 0, ps))
	//	//items = make([]ItemScore, 0, 100)
	//	//if e = redigo.ScanSlice(values, &items); e != nil {
	return redigo.Values(scon.Do("zrevrangebyscore", key, max, min, "withscores", "limit", 0, limit))
}

//
//移除有序集 key 中，所有 rank 值介于 min 和 max 之间(包括等于 min 或 max )的成员
func (rp *SsdbPool) ZRemRangeByRank(key interface{}, min, max int64) (interface{}, error) {
	conn := rp.getWrite()
	defer conn.Close()
	return conn.Do("zremrangebyrank", key, min, max)
}

//移除有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员
func (rp *SsdbPool) ZRemRangeByScore(key interface{}, min, max int64) (interface{}, error) {
	conn := rp.getWrite()
	defer conn.Close()
	return conn.Do("zremrangebyscore", key, min, max)
}
