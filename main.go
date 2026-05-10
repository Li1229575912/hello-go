package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello! This is v2.0.0. Running on %s/%s\n", runtime.GOOS, runtime.GOARCH)
	//Hello! This is v2.0.0. Running on darwin/arm64
}
