// Package src coding=utf-8
// @Project : knowledge-forge
// @Time    : 2025/9/11 10:55
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : parallel_merge.go
// @Software: GoLand
package src

// findCutPoint 在数组中找到分割点
func findCutPoint(X, Y []int, k int) (int, int) {
	lenx, leny := len(X), len(Y)
	i, j := 0, 0
	for k > 0 {
		if i == lenx {
			return i, j + k
		} else if j == leny {
			return i + k, j
		}

		// 二分找位置
		mid := k / 2
		newI := min(i+mid, lenx-1)
		newJ := min(j+mid, leny-1)

		if X[newI] <= Y[newJ] {
			k = k - (newI + 1 - i)
			i = newI + 1
		} else {
			k = k - (newJ + 1 - j)
			j = newJ + 1
		}
	}
	return i, j
}

func ParallelMerge(A, B []int) []int {
	lenA, lenB := len(A), len(B)
	result := make([]int, 0, lenA+lenB)

	if lenA == 0 {
		return append(result, B...)
	}
	if lenB == 0 {
		return append(result, A...)
	}

	mid := (lenA + lenB) / (2)
	i, j := findCutPoint(A, B, mid)

	// 两个协程分别计算
	var leftPart, rightPart []int
	parallelsWorkers := make(chan bool, 2)
	go func() {
		leftPart = ParallelMerge(A[:i], B[:j])
		parallelsWorkers <- true
	}()
	go func() {
		rightPart = ParallelMerge(A[i:], B[j:])
		parallelsWorkers <- true
	}()

	<-parallelsWorkers
	<-parallelsWorkers

	result = append(result, leftPart...)
	result = append(result, rightPart...)

	return result
}
