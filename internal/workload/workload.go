package workload

import (
	"time"
)

// CPUHeavy имитирует загрузку процессора
func CPUHeavy(n int, iter int) int {
	acc := 0
	for j := 0; j < iter; j++ {
		for i := 0; i < n; i++ {
			acc += i * i % (n + 1)
		}
	}
	return acc
}

// Allocate создаёт count объектов по size байт
func Allocate(count int, size int) [][]byte {
	data := make([][]byte, 0, count)
	for i := 0; i < count; i++ {
		buf := make([]byte, size)
		buf[0] = byte(i % 256) // трогаем, чтобы не оптимизировалось
		data = append(data, buf)
	}
	return data
}

// Sleep имитирует блокировку
func Sleep(d time.Duration) {
	time.Sleep(d)
}
