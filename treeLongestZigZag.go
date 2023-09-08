package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
const (
    start = iota
    left
    right
)


func Solution(T *Tree) int {
    // Implement your solution here
    return recZigzag(T, 0, start)
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func recZigzag (t *Tree, zigzag, turn int) int {
    max := zigzag

    if t.L != nil {
        currentMax := zigzag
        if turn == right {
            currentMax++
        }
        max = maxInt(recZigzag(t.L, currentMax, left), max)
    }

    if t.R != nil {
        currentMax := zigzag
        if turn == left {
            currentMax++
        }
        max = maxInt(recZigzag(t.R, currentMax, right), max)
    }
    return max
}