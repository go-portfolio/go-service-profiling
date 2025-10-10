package profiling

import (
	"os"
	"runtime/pprof"
)

// WriteThreadCreateProfile сохраняет профиль создания потоков (threadcreate)
func WriteThreadCreateProfile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return pprof.Lookup("threadcreate").WriteTo(f, 0)
}

// WriteMutexProfile сохраняет mutex профиль (конкуренция за мьютексы)
func WriteMutexProfile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return pprof.Lookup("mutex").WriteTo(f, 0)
}
