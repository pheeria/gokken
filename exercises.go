package gokken

func RecursiveSum(sequence []int) int {
	if len(sequence) == 1 {
		return sequence[0]
	}
	return sequence[0] + RecursiveSum(sequence[1:])
}

func RecursiveCount(sequence []int) int {
	if len(sequence) == 1 {
		return 1
	}
	return 1 + RecursiveCount(sequence[1:])
}

func RecursiveMax(sequence []int, max int) int {
	if max < sequence[0] {
		max = sequence[0]
	}
	if len(sequence) == 1 {
		return max
	}
	return RecursiveMax(sequence[1:], max)
}
