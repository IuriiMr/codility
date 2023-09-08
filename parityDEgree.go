package solution

// you can also use imports, for example:
// import "fmt"
import "math/bits"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
    i := bits.Len(uint(N))
	for i >= 0 {
		if (N % (1 << i)) == 0 {
			return i
		}
		i--
	}
	return 0

}