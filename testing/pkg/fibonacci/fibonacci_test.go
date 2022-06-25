package fibonacci

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

var fibonacciTestCases = []struct {
	in  uint
	out uint
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
	{11, 89},
	{12, 144},
	{13, 233},
	{14, 377},
	{15, 610},
	{16, 987},
	{17, 1597},
	{18, 2584},
	{19, 4181},
	{20, 6765},
}

func getFuncName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func TestFibonacci(t *testing.T) {
	for _, f := range []func(uint) uint{FibonacciSequential, FibonacciRecursive, FibonacciRecursiveWithCache} {
		fibonacciFunc := f
		funcName := strings.Split(getFuncName(fibonacciFunc), ".")[2]
		t.Run(funcName, func(t *testing.T) {
			t.Parallel()
			// t.Log(strings.Split(getFuncName(fibonacciFunc), ".")[2])
			for _, testCase := range fibonacciTestCases {
				if res := fibonacciFunc(uint(testCase.in)); res != testCase.out {
					t.Fatalf("incorrect output for %d: expected %d but got %d", testCase.in, testCase.out, res)
				}
			}
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for _, f := range []func(uint) uint{FibonacciSequential, FibonacciRecursive, FibonacciRecursiveWithCache} {
		fibonacciFunc := f
		funcName := strings.Split(getFuncName(fibonacciFunc), ".")[2]
		b.Run(funcName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, testCase := range fibonacciTestCases {
					if res := fibonacciFunc(uint(testCase.in)); res != testCase.out {
						b.Fatalf("incorrect output for %d: expected %d but got %d", testCase.in, testCase.out, res)
					}
				}
			}
		})
	}
}
