package algorithm

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

// 单元测试
func TestSimpleRandList(t *testing.T) {
	count,amount := int64(10), int64(100)
	randList := SimpleRandList(count,amount)
	SumFloat64List(amount,randList,t)
}

// 基准测试
func BenchmarkSimpleRandList(b *testing.B) {
	count,amount := int64(10), int64(100)
	for i := 0; i < b.N; i++ {
		SimpleRandList(count,amount)
	}
}

// RunParallel 测试并发性能
func BenchmarkSimpleRandList2(b *testing.B) {
	count,amount := int64(10), int64(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SimpleRandList(count,amount)
		}
	})
}

func TestAfterShuffle(t *testing.T) {
	count,amount := int64(10), int64(100)
	randList := AfterShuffle(count,amount)
	SumFloat64List(amount,randList,t)
}

func TestDoubleRandList(t *testing.T) {
	count,amount := int64(10), int64(100)
	randList := DoubleRandList(count,amount)
	SumFloat64List(amount,randList,t)
}

func TestDoubleAverageList(t *testing.T) {
	count,amount := int64(10), int64(100)
	randList := DoubleAverageList(count,amount)
	SumFloat64List(amount,randList,t)
}

func SumFloat64List(amount int64,randList []float64,t *testing.T) {
	var sum float64
	decimalSum := decimal.NewFromFloat(sum)
	for _,item := range randList {
		decimalItem := decimal.NewFromFloat(item)
		decimalSum = decimalSum.Add(decimalItem)
	}
	sum,_ = decimalSum.Float64()
	if sum != float64(amount) {
		fmt.Println(sum)
		t.Error("结果不正确")
	}
}