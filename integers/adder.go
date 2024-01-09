package integers

import "fmt"

// Add two integers
func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Printf("Add %d + %d = %d", 2, 2, Add(2, 2))
}
