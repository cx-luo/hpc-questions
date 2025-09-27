// Package src coding=utf-8
// @Project : knowledge-forge
// @Time    : 2025/9/11 10:57
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : happy_days.go
// @Software: GoLand
package src

var WeatherOptions = []string{"sunny", "rainy", "cloudy"}
var moodOptions = []string{"happy", "unhappy", "unknown"}

// SplitDayWeathers 分块处理，拆分成只有一种weather的子列表
//func SplitDayWeathers(w []string) [][]string {
//	var result [][]string
//	if len(w) == 0 {
//		return result
//	}
//
//	for i := 0; i < len(w); {
//		j := i
//		for j < len(w) && w[j] == w[i] {
//			j++
//		}
//		result = append(result, w[i:j])
//		i = j
//	}
//
//	return result
//}

func getMood(weather string, preMood string) string {
	switch weather {
	case "sunny":
		return "happy"
	case "rainy":
		return "unhappy"
	case "cloudy":
		return preMood
	default:
		return "unknown" // unknown
	}
}

func parallelCount(weather []string, defaultMood string) ([]string, string) {
	if len(weather) == 0 {
		return []string{defaultMood}, defaultMood
	}

	if len(weather) == 1 {
		mood := getMood(weather[0], defaultMood)
		return []string{mood}, mood
	}

	mid := len(weather) / 2
	var lastMood string

	leftPart := weather[:mid]
	rightPart := weather[mid:]

	var leftResult, rightResult []string

	// 处理左和右，都会影响上次结果的 lastMood
	leftResult, lastMood = parallelCount(leftPart, defaultMood)
	rightResult, lastMood = parallelCount(rightPart, lastMood)

	result := append(leftResult, rightResult...)

	return result, lastMood
}

// HappyDays 并行计算每日心情
func HappyDays(weather []string) []string {
	result, _ := parallelCount(weather, moodOptions[0])
	return result
}
