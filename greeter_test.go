package go_unit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// Assert
func TestStartGreeterAssert(t *testing.T) {
	result := StartGreeter("John")

	assert.Equal(t, "Hello John", result, "Result must be `Hello John`")
}

// Require
func TestStartGreeterRequire(t *testing.T) {
	result := StartGreeter("John")

	require.Equal(t, "Hello John", result, "Result must be 'Hello John'")
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on mac os")
	}

	result := StartGreeter("John")

	require.Equal(t, "Hello John", result, "Result must be 'Hello John'")
}

// Before and After (Test Main only run per package)
func TestMain(m *testing.M) {
	fmt.Println("Before")
	m.Run()
	fmt.Println("After")
}

// Sub Test
func TestSubTest(t *testing.T) {
	t.Run("John", func(t *testing.T) {
		result := StartGreeter("John")

		require.Equal(t, "Hello John", result, "Result must be 'Hello John'")
	})

	t.Run("Doe", func(t *testing.T) {
		result := StartGreeter("Doe")

		require.Equal(t, "Hello Doe", result, "Result must be 'Hello Doe'")
	})
}

// Test Table
func TestTableStartGreeter(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "StartGreeter(John)",
			request:  "John",
			expected: "Hello John",
		},
		{
			name:     "StartGreeter(Doe)",
			request:  "Doe",
			expected: "Hello Doe",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StartGreeter(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Benchmark
func BenchmarkStartGreeterJohn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StartGreeter("John")
	}
}

func BenchmarkStartGreeterDoe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StartGreeter("John")
	}
}

func BenchmarkStartGreeter(b *testing.B) {
	b.Run("John", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StartGreeter("John")
		}
	})

	b.Run("Doe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StartGreeter("Doe")
		}
	})
}

func BenchmarkStartGreeterTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "StartGreeter(John)",
			request: "John",
		},
		{
			name:    "StartGreeter(Doe)",
			request: "Doe",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StartGreeter(benchmark.request)
			}
		})
	}
}
