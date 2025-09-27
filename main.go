// Package xuhao_questions coding=utf-8
// @Project : knowledge-forge
// @Time    : 2025/9/10 17:25
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : main.go.go
// @Software: GoLand
package main

import (
	"fmt"
	"math/rand"
	"xuhao_questions/src"
)

func testParallelMerge() {
	A := []int{0, 1, 3, 5, 7, 9}
	B := []int{2, 4, 6, 8}

	mergedArray := src.ParallelMerge(A, B)
	fmt.Println("Merged array:", mergedArray)
}

func testHappyDays() {
	n := 10
	w := make([]string, n)
	//w[0] = "sunny"
	for i := 0; i < n; i++ {
		w[i] = src.WeatherOptions[rand.Intn(len(src.WeatherOptions))]
	}
	//w := []string{"sunny", "rainy", "sunny", "cloudy", "cloudy", "sunny", "rainy", "rainy", "rainy", "rainy"}
	//a := src.SplitDayWeathers(w)
	//fmt.Println(a)

	m := src.HappyDays(w)

	fmt.Println("weather array:\t", w)
	fmt.Println("Mood array:\t", m)
}

func testAdaptiveParallelSort() {
	data := make([]int, 50000)
	for i := range data {
		data[i] = rand.Intn(50000)
	}

	src.AdaptiveParallelSort(data)

	// 验证是否有序
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			fmt.Println(fmt.Sprintf("idx: {%d}, data: {%d}:{%d}", i, data[i], data[i-1]))
			panic("排序失败")
		}
	}
	fmt.Println("排序成功")
}

// 示例使用
func main() {
	testParallelMerge()
	testHappyDays()
	testAdaptiveParallelSort()
}
