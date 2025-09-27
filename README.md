# HPC Questions

A Go implementation of parallel algorithms and data processing techniques, featuring optimized sorting, merging, and weather-based mood calculation algorithms.

## Overview

This project demonstrates various parallel programming techniques in Go, including:

- **Parallel Merge Algorithm**: Efficiently merges two sorted arrays using concurrent processing
- **Happy Days Algorithm**: Calculates daily moods based on weather patterns using parallel processing
- **Adaptive Parallel QuickSort**: An optimized parallel sorting algorithm with adaptive thresholds

## Features

### 1. Parallel Merge (`parallel_merge.go`)
- Merges two sorted arrays in parallel using goroutines
- Uses binary search to find optimal cut points for parallel processing
- Handles edge cases for empty arrays efficiently

### 2. Happy Days (`happy_days.go`)
- Calculates daily moods based on weather conditions
- Supports three weather types: `sunny`, `rainy`, `cloudy`
- Uses parallel processing to handle large weather datasets
- Mood rules:
  - `sunny` → `happy`
  - `rainy` → `unhappy`
  - `cloudy` → maintains previous mood

### 3. Adaptive Parallel QuickSort (`optimized_quicksort.go`)
- Implements an adaptive parallel sorting algorithm
- Uses sampling technique to select optimal pivots
- Automatically switches to serial sorting for small arrays (threshold: 2000 elements)
- Configurable concurrency level (default: 100 goroutines)
- Includes custom semaphore implementation for concurrency control

### 4. Utilities (`utils.go`)
- Custom semaphore implementation for managing concurrent operations
- Provides `Acquire()`, `Release()`, and `Wait()` methods for goroutine synchronization

## Project Structure

```
hpc-questions/
├── main.go                    # Main entry point with test functions
├── go.mod                     # Go module definition
├── src/                       # Source code directory
│   ├── parallel_merge.go      # Parallel merge implementation
│   ├── happy_days.go          # Weather-based mood calculation
│   ├── optimized_quicksort.go # Adaptive parallel sorting
│   └── utils.go               # Utility functions and semaphore
└── README.md                  # This file
```

## Requirements

- Go 1.24.6 or later
- No external dependencies

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd hpc-questions
```

2. Build the project:
```bash
go build -o hpc-questions main.go
```

## Usage

### Running the Examples

Execute the main program to see all algorithms in action:

```bash
go run main.go
```

This will run three test functions:
- `testParallelMerge()`: Demonstrates parallel array merging
- `testHappyDays()`: Shows weather-based mood calculation
- `testAdaptiveParallelSort()`: Tests the parallel sorting algorithm

### Individual Algorithm Usage

#### Parallel Merge
```go
A := []int{0, 1, 3, 5, 7, 9}
B := []int{2, 4, 6, 8}
mergedArray := src.ParallelMerge(A, B)
```

#### Happy Days
```go
weather := []string{"sunny", "rainy", "cloudy", "sunny"}
moods := src.HappyDays(weather)
```

#### Adaptive Parallel Sort
```go
data := []int{5, 2, 8, 1, 9, 3}
sorted := src.AdaptiveParallelSort(data)
```

## Algorithm Details

### Parallel Merge Algorithm
- **Time Complexity**: O(log n) for finding cut points, O(n) for merging
- **Space Complexity**: O(n)
- **Concurrency**: Uses 2 goroutines for parallel processing

### Happy Days Algorithm
- **Time Complexity**: O(n) with parallel processing
- **Space Complexity**: O(n)
- **Concurrency**: Recursive parallel processing with mood state propagation

### Adaptive Parallel QuickSort
- **Time Complexity**: O(n log n) average case
- **Space Complexity**: O(n)
- **Concurrency**: Configurable (default: 100 goroutines)
- **Adaptive**: Switches to serial sorting for arrays < 2000 elements

## Configuration

You can modify the following constants in `optimized_quicksort.go`:

```go
const (
    NumConcurrency = 100  // Number of concurrent goroutines
    Threshold      = 2000 // Switch to serial sorting threshold
)
```

## Performance Considerations

- The parallel algorithms are most effective on large datasets
- Small arrays may perform better with serial algorithms due to goroutine overhead
- Memory usage scales with input size and concurrency level
- The adaptive threshold helps balance parallel efficiency with overhead

## Author

- **Author**: chengxiang.luo
- **Email**: chengxiang.luo@foxmail.com
- **Project**: knowledge-forge

## License

This project is part of the knowledge-forge project. Please refer to the main project for licensing information.

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests to improve the algorithms or add new parallel processing techniques.
