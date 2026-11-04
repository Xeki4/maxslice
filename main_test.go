package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	t.Run("нулевой размер", func(t *testing.T) {
		elements := generateRandomElements(0)
		if len(elements) != 0 {
			t.Errorf("generateRandomElements(0) должен вернуть пустой слайс, получено длинну %d", len(elements))
		}
	})

	t.Run("положительный размер", func(t *testing.T) {
		size := 100
		elements := generateRandomElements(size)
		if len(elements) != size {
			t.Errorf("generateRandomElements(%d) должен вернуть слайс длины %d, получено %d", size, size, len(elements))
		}
		// Все сгенерированные числа должны быть положительными.
		for i, value := range elements {
			if value <= 0 {
				t.Errorf("элемент [%d] = %d не является положительным числом", i, value)
			}
		}
	})
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"два элемента", []int{3, 7}, 7},
		{"максимум в начале", []int{100, 1, 2, 3}, 100},
		{"максимум в середине", []int{1, 50, 2, 3}, 50},
		{"максимум в конце", []int{1, 2, 3, 99}, 99},
		{"одинаковые элементы", []int{5, 5, 5}, 5},
		{"большие числа", []int{1_000_000, 2, 999_999}, 1_000_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.in); got != tt.want {
				t.Errorf("maximum(%v) = %d, хотелось %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"меньше 8 элементов", []int{3, 1, 4, 1, 5}, 5},
		{"ровно 8 элементов", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"больше 8 элементов", []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 100},
		{"максимум в начале", []int{100, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.in); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, хотелось %d", tt.in, got, tt.want)
			}
		})
	}
}
