package sort

func BubbleSort(sequence []int) {
    for i := 0; i < len(sequence); i++ {
        for j := 0; j < len(sequence) - 1 - i; j++ {
            if sequence[j] > sequence[j + 1] {
                swap(sequence, j, j + 1)
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
                swap(sequence, j, j + 1)
                hasSwapped = true
            }
        }
    }
}
