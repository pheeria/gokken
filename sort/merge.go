package sort

func merge(left, right []int) []int{
    var result []int

    i := 0
    j := 0

    for i < len(left) || j < len(right) {
        if i == len(left) {
            result = append(result, right[j:]...)
            break
        }
        if j == len(right) {
            result = append(result, left[i:]...)
            break
        }
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }

    return result
}

func MergeSort(sequence []int) []int {
    if len(sequence) == 1 {
        return sequence
    }

    middle := len(sequence) / 2
    left := MergeSort(sequence[:middle])
    right := MergeSort(sequence[middle:])

    return merge(left, right)
}
