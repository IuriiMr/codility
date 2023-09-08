package solution

// you can also use imports, for example:
// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // Implement your solution here
    l := len(A) + 1

    currentSum := 0
    for _, v := range A {
        currentSum += v
    }
    shouldBeSum := l * (l + 1) / 2
    
    return shouldBeSum - currentSum
}