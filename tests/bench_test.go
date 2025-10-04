package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/go-portfolio/go-service-profiling/internal/handlers"
)

func BenchmarkWorkHandler(b *testing.B) {
	req := httptest.NewRequest("GET", "/work", nil)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handlers.WorkHandler(w, req)
	}
}

func BenchmarkAllocHandler(b *testing.B) {
	req := httptest.NewRequest("GET", "/alloc", nil)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handlers.AllocHandler(w, req)
	}
}
