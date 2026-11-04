package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// CHUNKS — количество частей, на которые разбивается слайс для параллельного поиска максимума.
const CHUNKS = 8

// SIZE — размер сгенерированного слайса чисел.
const SIZE = 10_000_000

// generateRandomElements генерирует слайс заданного размера, заполненный случайными
// положительными целыми числами. Если передан размер, равный 0 (или negativo),
// возвращается пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := range elements {
		// rand.Intn(n) возвращает число из [0, n), прибавляем 1, чтобы число было положительным (> 0).
		elements[i] = rand.Intn(1_000_000_000) + 1
	}
	return elements
}

// maximum находит максимальное значение в слайсе.
// Обрабатывает крайние случаи: пустой слайс возвращает 0, слайс с одним элементом возвращает его.
func maximum(elements []int) int {
	if len(elements) == 0 {
		return 0
	}

	max := elements[0]
	for _, value := range elements[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks находит максимальное значение, разбивая слайс на CHUNKS частей и ища
// максимум в каждой части в отдельной горутине. Затем ищет максимум среди
// найденных значений. Используется sync.WaitGroup для ожидания горутин.
func maxChunks(elements []int) int {
	if len(elements) == 0 {
		return 0
	}

	chunkSize := len(elements) / CHUNKS
	maxes := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		// Последняя часть забирает остаток элементов, чтобы не потерять данные,
		// если длина слайса не делится на CHUNKS без остатка.
		if i == CHUNKS-1 {
			end = len(elements)
		}

		go func(slice []int, index int) {
			defer wg.Done()
			maxes[index] = maximum(slice)
		}(elements[start:end], i)
	}
	wg.Wait()

	return maximum(maxes)
}

func main() {
	elements := generateRandomElements(SIZE)

	// Поиск максимума в один поток.
	start := time.Now()
	singleMax := maximum(elements)
	singleTime := time.Since(start).Microseconds()

	// Поиск максимума в 8 потоков (горутинах).
	start = time.Now()
	chunkMax := maxChunks(elements)
	chunkTime := time.Since(start).Microseconds()

	fmt.Printf("Размер слайса: %d\n", SIZE)
	fmt.Printf("Максимум (1 поток):  %d, время: %d мкс\n", singleMax, singleTime)
	fmt.Printf("Максимум (8 потоков): %d, время: %d мкс\n", chunkMax, chunkTime)
	if singleMax == chunkMax {
		fmt.Println("Результаты совпадают.")
	} else {
		fmt.Println("ВНИМАНИЕ: результаты не совпадают!")
	}
}
