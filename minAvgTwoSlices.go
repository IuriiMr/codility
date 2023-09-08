package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func minForTwo (A []int) (float64, int) {
    minAvg := 10.0
    minStart := 0
    for i, v := range A {
        if i <= len(A) - 2 {
            avgTwo := float64(A[i + 1] + v) / 2.0
            if avgTwo < minAvg {
                minAvg = avgTwo
                minStart = i
            }
        }
    }
    return minAvg, minStart
}

func minForThree (A []int) (float64, int) {
    minAvg := 10.0
    minStart := 0
    for i, v := range A {
        if i <= len(A) - 3 {
            avgThree := float64(A[i + 1] + A[i + 2] + v) / 3.0
            if avgThree < minAvg {
                minAvg = avgThree
                minStart = i
            }
        }
    }
    return minAvg, minStart
}

func Solution(A []int) int {
    // Implement your solution here
    
    minAvgTwo, minStartTwo := minForTwo(A)
    minAvgThree, minstartThree := minForThree(A)
    
    if minAvgThree > minAvgTwo {
        return minStartTwo
    }
    if minAvgTwo > minAvgThree {
        return minstartThree
    }
    if minStartTwo <= minstartThree {
        return minStartTwo
    }
    return minstartThree
}