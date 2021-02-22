package algorithm

import (
	"math/rand"
	"time"
)

const MIN_AMOUNT = 1

// amount 单位是分
// SimpleRand 简单随机
func SimpleRand(count,amount int64) int64 {
	if count == 1 {
		return amount
	}
	max := amount - count *MIN_AMOUNT
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int63n(max)
	return randNum+ MIN_AMOUNT
}

// SimpleRandList 简单随机红包列表
func SimpleRandList(count,amount int64) []float64 {
	randList := make([]float64,0)
	amount = amount * 100
	for i := int64(0); i < count; i++ {
		randAmount := SimpleRand(count-i,amount)
		amount = amount - randAmount
		randList = append(randList, float64(randAmount)/float64(100))
	}
	return randList
}

// AfterShuffle 后洗牌算法
func AfterShuffle(count,amount int64) []float64 {
	randList := make([]float64,0)
	amount = amount * 100
	for i := int64(0); i<count; i++ {
		randAmount := SimpleRand(count-i,amount)
		amount = amount - randAmount
		randList = append(randList, float64(randAmount)/float64(100))
	}
	rand.Shuffle(len(randList),func(i, j int){
		randList[i],randList[j] = randList[j],randList[i]
	})
	return randList
}

// amount 单位是分
// DoubleShuffle 两次随机算法（先洗牌随机算法改进）
func DoubleShuffle(count,amount int64) int64 {
	if count == 1 {
		return amount
	}
	max := amount - MIN_AMOUNT*count
	rand.Seed(time.Now().UnixNano())
	seed := rand.Int63n(count*2) + 1
	n := max / seed + MIN_AMOUNT
	x := rand.Int63n(n)
	return x + MIN_AMOUNT
}

// DoubleRandList 两次随机算法
func DoubleRandList(count,amount int64) []float64 {
	randList := make([]float64,0)
	amount = amount * 100
	for i := int64(0); i < count; i++ {
		randNum := DoubleShuffle(count - i,amount)
		amount -= randNum
		randList = append(randList, float64(randNum)/float64(100))
	}
	return randList
}

// amount的单位是分
// DoubleAverage 二倍均值算法
func DoubleAverage(count,amount int64) int64 {
	if count == 1 {
		return amount
	}
	max := amount - count * MIN_AMOUNT
	avg := max / count
	doubleAvg := avg * 2 + 1
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int63n(doubleAvg)
	return randNum + MIN_AMOUNT
}

// DoubleAverageList 二倍均值算法
func DoubleAverageList(count,amount int64) []float64 {
	randList := make([]float64,0)
	amount = amount * 100
	for i:= int64(0); i < count; i++ {
		randNum := DoubleAverage(count - i, amount)
		amount -= randNum
		randList = append(randList, float64(randNum) / float64(100))
	}
	return randList
}