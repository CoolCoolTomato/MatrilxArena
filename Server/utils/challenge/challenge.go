package challenge

import "math"

// DynamicScore 动态计算分数
// s 初始分数
// r 最低分数
// d 衰减速度
// x 解题率
func DynamicScore(s float64, r float64, d float64, x float64) int {
	exponentialPart := math.Exp((1 - x) / d)
	result := s * (r + (1-r)*exponentialPart)
	return int(result)
}
