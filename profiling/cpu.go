package profiling

import (
	"os"
	"runtime/pprof"
)

// StartCPUProfile запускает запись CPU профиля в файл.
func StartCPUProfile(filename string) (*os.File, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		f.Close()
		return nil, err
	}
	return f, nil
}

// StopCPUProfile останавливает CPU профилирование.
func StopCPUProfile(f *os.File) {
	pprof.StopCPUProfile()
	if f != nil {
		f.Close()
	}
}
