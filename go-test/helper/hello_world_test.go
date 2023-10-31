package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Muhammad")
	}
}

func BenchmarkHelloWorldRama(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ramadhan")
	}
}
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Muhammad",
			request:  "Muhammad",
			expected: "Hello Muhammad",
		},
		{
			name:     "Ramadhan",
			request:  "Ramadhan",
			expected: "Hello Ramadhan",
		},
	}

	for _, tests := range tests {
		t.Run(tests.name, func(t *testing.T) {
			result := HelloWorld(tests.request)
			require.Equal(t, tests.expected, result)
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Muhammad", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Muhammad")
		}
	})

	b.Run("Ramadhan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ramadhan")
		}
	})
}

// func TestSubTest(t *testing.T) {
// 	t.Run("Rama", func(t *testing.T) {
// 		result := HelloWorld("Rama")
// 		assert.Equal(t, "Hello Rama", result, "Result must be 'Hello Ramadhan'")
// 	})
// 	t.Run("Ramadhan", func(t *testing.T) {
// 		result := HelloWorld("Ramadhan")
// 		assert.Equal(t, "Hi Ramadhan", result, "Result must be 'Hello Ramadhan'")
// 	})
// }

// // TestMain
// func TestMain(m *testing.M) {
// 	// before
// 	fmt.Println("Before unit test")

// 	m.Run()

// 	// after
// 	fmt.Println("After unit test")
// }

// // TestSkip
// func TestSkip(t *testing.T) {
// 	if runtime.GOOS == "linux" {
// 		t.Skip("Cant run on Mac OS")
// 	}
// 	result := HelloWorld("Ramadhan")
// 	assert.Equal(t, "Hello Ramadhan", result, "Result must be 'Hello Ramadhan'")
// 	fmt.Println("Test Done")
// }

// // Assert & Require
// func TestHelloWorldAssert(t *testing.T) {
// 	result := HelloWorld("Ramadhan")
// 	assert.Equal(t, "Hello Ramadhan", result, "Result must be 'Hello Ramadhan'")
// 	fmt.Println("Test Done")
// }

// func TestHelloWorldRequire(t *testing.T) {
// 	result := HelloWorld("Ramadhan")
// 	assert.Equal(t, "Hello Ramadhan", result, "Result must be 'Hello Ramadhan'")
// 	fmt.Println("Test Done")
// }

// func TestHelloWorld(t *testing.T) {
// 	result := HelloWorld("Rama")

// 	if result != "Hello Rama" {
// 		t.Error("Result must be 'Hello Rama'")
// 	}

// 	fmt.Println("Berhasil")
// }

// func TestHelloWorldRama(t *testing.T) {
// 	result := HelloWorld("Ramadhan")

// 	if result != "Hi Ramadhan" {
// 		t.Error("Result must be 'Hello Ramadhan'")
// 	}

// 	fmt.Println("Berhasil")
// }
