package gokken

func QuickSort(sequence []int) []int {
    if len(sequence) < 2 {
        return sequence
    }

    pivot := sequence[0]
    var less []int
    var greater []int

    for _, v := range sequence[1:] {
        if v <= pivot {
            less = append(less, v)
        } else {
            greater = append(greater, v)
        }
    }

    return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}
