// Package src coding=utf-8
// @Project : knowledge-forge
// @Time    : 2025/9/11 18:55
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : optimized_quicksort.go
// @Software: GoLand
package src

import (
	"math/rand"
	"sort"
	"sync"
)

/*
在lecture 3的例子中，可以看出pivot的选择对排序有很大影响，
选择不好容易造成左右子树不平衡，甚至退化为单支树结构（depth: O(logN)）。

我的优化思路：
就像处理数据库中大批量数据在不同节点中，因分区策略不完善造成数据分布不均衡的问题，
可以采用的一种方式是增加二级分区，这样并行时可以提升效率。
同事，在数据量不达到一定数量级时，可以使用普通的快排或者其他排序算法。

function optimizedParallelQuickSort(A, low, high)
    // low, high 是数组起始index
    if high - low < threshold then
        // 对于小数组，直接使用串行快速排序
        quickSort(A, low, high) or mergeSort(A, low, high)
    else
        secondPartitionParallelSort(A, low, high)

function secondPartitionParallelSort(A, low, high)
    n = high - low + 1
    samples, pivots = []
    NumConcurrency, sampleSize int   // NumConcurrency 并行度

    // 采样，用来选取pivot
    for i = 0; i < sampleSize; i++:
        idx = low + random() % n    // 随机再hash
        samples.append(A[idx])

    // 排序，根据并行度选取pivots
    sort(samples)
    for j = 1 to NumConcurrency -1:
        pos = j * len(samples) / NumConcurrency
        pivots.append(samples[pos])

    sub_partitions = init(array) * NumConcurrency
    // 元素分组
    parallel_for i = low to high do:
        value = A[i]
        partition_id = findPartitionId(value, pivots)
        sub_partitions[partition_id].add(value)

    // 并行sub_partition排序
    parallel_for sub_partition in sub_partitions:
        sorted_sub_partition = optimizedParallelQuickSort(sub_partition, 0, len(sub_partition)-1)
        sub_partition.inPlaceReplace(sorted_sub_partition) // 原地更新

    // 写回原数组，串行写，依赖前一个offset
    writeOffset = low
    for sub_partition in sub_partitions:
        copy(A[writeOffset, writeOffset+len(sub_partition)-1])
        writeOffset += len(sub_partition)
*/

const (
	NumConcurrency = 100
	Threshold      = 2000
)

var SampleSize = 100

// 自己写的一个控制并发度的工具
var sema = NewSemaphore(NumConcurrency)

func optimizedParallelQuickSort(A []int, low, high int) []int {
	if high-low < Threshold {
		return quickSort(A, low, high)
	} else {
		return secondPartitionParallelSort(A, low, high)
	}
}

func quickSort(A []int, low, high int) []int {
	// 基准情况：区间无效或太小
	if low >= high {
		return A
	}
	if high-low < Threshold {
		sort.Ints(A[low : high+1])
		return A
	}

	mid := low + (high-low)/2
	pivot := A[mid]

	// 分区：将数组分为 <= pivot 和 > pivot 两部分
	left, right := low, high
	for left <= right {
		for A[left] < pivot {
			left++
		}
		for A[right] > pivot {
			right--
		}
		if left <= right {
			A[left], A[right] = A[right], A[left]
			left++
			right--
		}
	}

	quickSort(A, low, right)
	quickSort(A, left, high)

	return A
}

func secondPartitionParallelSort(A []int, low, high int) []int {
	n := high - low + 1
	var pivots, samples []int

	if n < SampleSize {
		SampleSize = n
	}

	for i := 0; i < SampleSize; i++ {
		idx := low + rand.Intn(n) // 随机生成 [0, n) 的整数
		samples = append(samples, A[idx])
	}
	sort.Ints(samples)

	if len(samples) >= 2 {
		pivots = make([]int, 0, NumConcurrency-1)
		for i := 1; i < NumConcurrency; i++ {
			idx := i * len(samples) / NumConcurrency
			val := samples[idx]
			pivots = append(pivots, val)
		}
	}

	subPartitions := make([][]int, NumConcurrency)
	var mu sync.Mutex
	for i := low; i <= high; i++ {
		v := A[i]
		sema.Acquire(1)
		go func(value int) {
			defer sema.Release()
			partitionId := findPartitionId(value, pivots)

			// 加锁，避免并发 append
			mu.Lock()
			subPartitions[partitionId] = append(subPartitions[partitionId], value)
			mu.Unlock()
		}(v)
	}
	sema.Wait()

	var sortWg = NewSemaphore(NumConcurrency)
	for i, partition := range subPartitions {
		if len(partition) == 0 {
			continue
		}
		sortWg.Acquire(1)
		go func(idx int, p []int) {
			defer sortWg.Release()
			optimizedParallelQuickSort(subPartitions[idx], 0, len(subPartitions[idx])-1)
		}(i, partition)
	}
	sortWg.Wait()

	writeOffset := low
	for _, partition := range subPartitions {
		copy(A[writeOffset:], partition)
		writeOffset += len(partition)
	}
	return A
}

func findPartitionId(v int, pivots []int) int {
	for i, pivot := range pivots {
		if v < pivot {
			return i
		}
	}

	return len(pivots)
}

func AdaptiveParallelSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	return optimizedParallelQuickSort(a, 0, len(a)-1)
}
