func factorial(num int) int {
    result := 1
    for ; num > 0; num-- {
        result *= num
    }
    return result
}

func main() {
    fmt.Println(factorial(20)) // 2432902008176640000
}

