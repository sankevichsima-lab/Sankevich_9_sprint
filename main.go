package main

import (
	"fmt"
	"math/rand"
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

	maxData := data[0]
	for _, val := range data[1:] {
		if val > maxData {
			maxData = val
		}
	}
	return maxData
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}

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

			topMax[idx] = maximum(data[s:e])
		}(i, start, end)
	}
	wg.Wait()
	return maximum(topMax)
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
