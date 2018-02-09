package main // for all executable files

import (
	"fmt"           // formatting
	"hello/lesson1" // custom package
	"os"
	"runtime"
)

// ===cmd===:
// go build
// go run script.go

// entry point
func main() {
	fmt.Println("Hello world! From", runtime.GOOS)

	// Constants:
	const a = 123
	fmt.Println("Const a=", a)

	fmt.Println("Current user:", os.Getenv("USERNAME"))

	lesson1.Varentry()
}
