package fibonacci

func FibonacciSequential(n uint) uint {
	if n < 2 {
		return n
	}

	prev2, prev1 := 0, 1
	var i uint
	for i = 2; i < n; i++ {
		prev2, prev1 = prev1, prev2+prev1
	}

	return uint(prev2 + prev1)
}

func FibonacciRecursive(n uint) uint {
	if n < 2 {
		return n
	}
	return FibonacciRecursive(n-2) + FibonacciRecursive(n-1)
}

var (
	fibsCache = make(map[uint]uint)
)

func FibonacciRecursiveWithCache(n uint) uint {
	if n < 2 {
		return n
	}

	var res uint
	if v, ok := fibsCache[n]; ok {
		res = v
	} else {
		res = FibonacciRecursiveWithCache(n-1) + FibonacciRecursiveWithCache(n-2)
		fibsCache[n] = res
	}
	return res
}
