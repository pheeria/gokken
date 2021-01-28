package search

func IterativeBinarySearch(list []int, target int) int {
    low := 0
    high := len(list) - 1

    for low <= high {
        mid := (high + low) / 2
        if list[mid] > target {
            high = mid - 1
        } else if list[mid] < target {
            low = mid + 1
        } else {
            return mid
        }
    }

    return -1
}

func RecursiveBinarySearch(list []int, target int) int {
    return recursive(list, 0, len(list) - 1, target)
}

func recursive(list []int, low int, high int, target int) int {
    if low > high {
        return - 1
    }

    mid := (high + low) / 2
    if list[mid] == target {
        return mid
    }

    if list[mid] > target {
        return recursive(list, low, mid - 1, target)
    } else if list[mid] < target {
        return recursive(list, mid + 1, high, target)
    }

    return -1
}
