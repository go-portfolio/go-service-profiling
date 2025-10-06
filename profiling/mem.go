package profiling

import (
	"os"
	"runtime"
	"runtime/pprof"
)

// WriteHeapProfile сохраняет heap-профиль (используемая память).
func WriteHeapProfile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	runtime.GC() // запускаем сборку мусора для точности
	return pprof.WriteHeapProfile(f)
}

// WriteAllocsProfile сохраняет профиль распределений (allocs).
func WriteAllocsProfile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return pprof.Lookup("allocs").WriteTo(f, 0)
}
