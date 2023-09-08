package solution

// you can also use imports, for example:
// import "fmt"
// import "os"


// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
    // Implement your solution here

	last := -1
	max := 1

	for i := 0; i < 32; i++ {
		if (N & (1 << i)) != 0 {
			if last >= 0 {
				if i - last > max {
					max = i - last
				}
			}
			last = i
		}
	}
	return max - 1
}

