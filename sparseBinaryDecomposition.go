package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func isSparce (n int) bool {
	for i := 0; i < 29; i++ {
		if (n & (1 << i)) != 0 && (n & (1 << (i + 1))) != 0 {
			return false
		}
	}
	return true
}

func Solution(N int) int {
    // Implement your solution here
	ans := N & 0x55555555
	if isSparce(N - ans) {
		return ans
	} 
	ans = N & 0xAAAAAAAA
	if isSparce(N - ans) {
		return ans
	}
	return -1
}