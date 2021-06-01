package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkTableHelloWorld(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "david benchmark",
			request: "david",
		},
		{
			name:    "ayu benchmark",
			request: "ayu",
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

func BenchmarkSubHelloWorld(b *testing.B) {
	b.Run("David", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("David")
		}
	})

	b.Run("Ayu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ayu")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("David")
	}
}

func TestHelloWorldTable(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "David",
			args: args{
				name: "david",
			},
			want: "david",
		},
		{
			name: "Ayu",
			args: args{
				name: "ayu",
			},
			want: "ayu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HelloWorld(tt.args.name); got != tt.want {
				t.Errorf("HelloWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubFuncHelloWorld(t *testing.T) {
	t.Run("david", func(t *testing.T) {
		result := HelloWorld("david")
		if result != "hello david" {
			// fail unit test
			t.Error("result is not hello david")
		}
	})

	t.Run("david 1", func(t *testing.T) {
		result := HelloWorld("david 1")
		if result != "hello david 1" {
			// fail unit test
			t.Error("result is not hello david 1")
		}
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Run test")
	m.Run()
	fmt.Println("After Run test")
}

func TestSkipNotMatchOS(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("cannot run in linux")
	}

	result := HelloWorld("david")
	if result != "hello david" {
		// fail unit test
		t.Error("result is not hello david")
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("david")
	if result != "hello david" {
		// fail unit test
		t.Error("result is not hello david")
	}

	fmt.Println("test hello world done")
}
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("david")
	assert.Equal(t, "hello david", result, "result must be 'hello david' ")
	fmt.Println("execute finished")
}
