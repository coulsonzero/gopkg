package pro

import (
	"math/rand"
	"time"
)

// 最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 最大公约数
func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

// 最小公倍数
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// 生成随机数 (1, n)之间
func randInt(n int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(n)
}
