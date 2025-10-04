package workload

import (
	"time"
)

// CPUHeavy имитирует нагрузку на процессор
// n — количество итераций в одном цикле
// iter — количество внешних повторений
// возвращает аккумулятор, чтобы компилятор не оптимизировал цикл
func CPUHeavy(n int, iter int) int {
	acc := 0
	for j := 0; j < iter; j++ {       // внешний цикл повторений
		for i := 0; i < n; i++ {      // внутренний цикл вычислений
			acc += i * i % (n + 1)    // простая математика для нагрузки CPU
		}
	}
	return acc
}

// Allocate создаёт count объектов по size байт
// Используется для имитации потребления памяти
func Allocate(count int, size int) [][]byte {
	data := make([][]byte, 0, count) // создаём слайс для хранения объектов
	for i := 0; i < count; i++ {
		buf := make([]byte, size)      // создаём объект размером size байт
		buf[0] = byte(i % 256)         // обращаемся к первому элементу, чтобы память реально выделилась
		data = append(data, buf)       // добавляем объект в слайс
	}
	return data
}

// Sleep имитирует блокировку или задержку
// d — длительность сна
func Sleep(d time.Duration) {
	time.Sleep(d) // "усыпляем" горутину на заданное время
}
