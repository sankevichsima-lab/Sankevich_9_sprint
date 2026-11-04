package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size == 0 {
		return nil
	}
	mas := make([]int, size)
	for i := 0; i < size; i++ {
		mas[i] = rand.Int()
	}
	return mas
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	res := slices.Max(data)
	return res
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	topMax := make([]int, CHUNKS)
	var wg sync.WaitGroup

	size := len(data) / CHUNKS
	for i := 0; i < CHUNKS; i++ {
		start := i * size
		end := (i + 1) * size
		if i == CHUNKS-1 {
			end = len(data)
		}
		wg.Add(1)
		go func(idx, s, e int) {
			defer wg.Done()

			topMax[idx] = slices.Max(data[s:e])
		}(i, start, end)
	}
	wg.Wait()
	result := slices.Max(topMax)
	return result
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	slice := make([]int, SIZE)
	slice = generateRandomElements(SIZE)

	startTime := time.Now()
	fmt.Println("Ищем максимальное значение в один поток")
	max := maximum(slice)
	elapsed := time.Since(startTime)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())

	startTime = time.Now()
	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	max = maxChunks(slice)
	elapsed = time.Since(startTime)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
