package profiling

import (
	"os"
	"runtime/trace"
)

// StartTrace запускает трассировку выполнения.
func StartTrace(filename string) (*os.File, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	if err := trace.Start(f); err != nil {
		f.Close()
		return nil, err
	}
	return f, nil
}

// StopTrace останавливает трассировку.
func StopTrace(f *os.File) {
	trace.Stop()
	if f != nil {
		f.Close()
	}
}
