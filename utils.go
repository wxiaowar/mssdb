package mssdb

// 生成redis分页
func buildRange(cur, ps int) (int, int) {
	begin := 0
	if cur > 1 {
		begin = (cur - 1) * ps
	}
	end := begin + ps - 1

	return begin, end
}

//生成score，时间会以秒数的形式存储在64位整型的前(64-bits)位，tag会存储在后bits位。
func MakeZScore(unix int64, tag uint32, bits uint) int64 {
	return (unix << bits) + int64(tag)
}

//从score中提取标签的值
func GetTagFromScore(score int64, bits uint) uint32 {
	return uint32(score & ((int64(1) << bits) - 1))
}

//从score中提取时间
func GetTimeFromScore(score int64, bits uint) int64 {
	return score>>bits
}
