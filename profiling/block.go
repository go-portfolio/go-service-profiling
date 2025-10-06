package profiling

import (
	"os"
	"runtime"
	"runtime/pprof"
)

// EnableBlockProfile включает профилирование блокировок.
// rate = число стэктрейсов блокировок в секунду (обычно 1).
func EnableBlockProfile(rate int) {
	runtime.SetBlockProfileRate(rate)
}

// WriteBlockProfile сохраняет block-профиль в файл.
func WriteBlockProfile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return pprof.Lookup("block").WriteTo(f, 0)
}
