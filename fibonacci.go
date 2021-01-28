package gokken

func IterativeFibonacci(size int) []int {
	result := make([]int, size)
	result[0] = 0
	result[1] = 1

	for i := 2; i < size; i++ {
		result[i] = result[i-1] + result[i-2]
	}

	return result
}

func RecursiveFibonacci(size int) []int {
	var fib func(int) int
	fib = func(number int) int {
		if number < 2 {
			return number
		}
		return fib(number-1) + fib(number-2)
	}

	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = fib(i)
	}
	return result
}

func ClosureFibonacci(size int) []int {
	var fib func() func() int
	fib = func() func() int {
		a := 0
		b := 1
		return func() int {
			tmp := a

			a = b
			b = tmp + b

			return tmp
		}
	}

	result := make([]int, size)
	fibonacci := fib()
	for i := 0; i < size; i++ {
		result[i] = fibonacci()
	}
	return result
}
