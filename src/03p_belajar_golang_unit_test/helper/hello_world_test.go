package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Budi",
			request: "Budi",
		},
		{
			name:    "Darmawan",
			request: "Darmawan",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Budi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Budi")
		}
	})
	b.Run("Darmawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Darmawan")
		}
	})
}

// Benchmark
func BenchmarkHelloWorldBudi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Budi")
	}
}

func BenchmarkHelloWorldDarmawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Darmawan")
	}
}

// Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Darmawan",
			request:  "Darmawan",
			expected: "Hello Darmawan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		require.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'")
	})
	t.Run("Darmawan", func(t *testing.T) {
		result := HelloWorld("Darmawan")
		require.Equal(t, "Hello Darmawan", result, "Result must be 'Hello Darmawan'")
	})
}

// Main Test (before dan after unit test)
func TestMain(m *testing.M) {
	// before unit test
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after unit test
	fmt.Println("AFTER UNIT TEST")
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := HelloWorld("Budi")
	require.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'")
}

// Cara lebih mudah menggunakan library testify/assert (recommended)
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Budi")
	assert.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'") // memanggil Fail()
	fmt.Println("TestHelloWorld with Assert Done")                       // akan dieksekusi jika ada error
}

// Cara lebih mudah menggunakan library testify/require (recommended)
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Budi")
	require.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'") // memanggil FailNow()
	fmt.Println("TestHelloWorld with Require Done")                       // tidak akan dieksekusi jika ada error
}

// Cara manual menggunakan if else (not recommended)
func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")

	if result != "Hello Budi" {
		// Error
		// panic("Result is not 'Hello Budi'") // program berhenti
		// t.Fail() // tidak memberikan informasi error, perintah selanjutnya MASIH dieksekusi
		t.Error("Result must be 'Hello Budi'")
	}

	fmt.Println("TestHelloWorldBudi Done")
}

// Cara manual menggunakan if else (not recommended)
func TestHelloWorldDarmawan(t *testing.T) {
	result := HelloWorld("Darmawan")

	if result != "Hello Darmawan" {
		// Error
		// panic("Result is not 'Hello Darmawan'") // program berhenti
		// t.FailNow() // tidak memberikan informasi error, perintah selanjutnya TIDAK dieksekusi
		t.Fatal("Result must be 'Hello Darmawan'")
	}
	fmt.Println("TestHelloWorldDarmawan Done")

}
