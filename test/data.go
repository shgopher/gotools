package test

import "github.com/googege/gotools/public"

func TestDatasSliceInt(n int) []int {
	result := make([]int, 0, n)
	r := public.NewRand()
	for i := 10; i <= n*10; i++ {
		result = append(result, r.Intn(20)+r.Intn(30))
	}
	return result
}
