package gokken

func BubbleSort(sequence []int) {
    for i := 0; i < len(sequence); i++ {
        for j := 0; j < len(sequence) - 1 - i; j++ {
            if sequence[j] > sequence[j + 1] {
                tmp := sequence[j]
                sequence[j] = sequence[j + 1]
                sequence[j + 1] = tmp
            }
        }
    }
}

func OptimizedBubbleSort(sequence []int) {
    hasSwapped := true
    for hasSwapped {
        hasSwapped = false
        for j := 0; j < len(sequence) - 1; j++ {
            if sequence[j] > sequence[j + 1] {
                tmp := sequence[j]
                sequence[j] = sequence[j + 1]
                sequence[j + 1] = tmp
                hasSwapped = true
            }
        }
    }
}
