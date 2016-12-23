// UtilFunctions
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomString(strlen int) string {
	rand.Seed(time.Now().UnixNano())
	arr := make([]byte, strlen)
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	for x := 0; x < strlen; x++ {
		arr[x] = chars[rand.Intn(len(chars))]
	}
	return string(arr)
}

func main() {
	fmt.Println("Hello World!")
}
