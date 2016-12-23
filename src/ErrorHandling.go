// ErrorHandling
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func fileIO() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("File not created and creating new one")
		file, err = os.Create("file.txt")
	} else {
		fmt.Println("File created", *file)
	}
	for x := 0; x < 100; x++ {
		size, err1 := file.WriteString(generateRandomString(x + 1))
		if err1 != nil {
			fmt.Println("Number of bytes written", size)
		}
	}

}

func fileIOWithFlags() {
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		size, _ := file.WriteString(generateRandomString(10))
		fmt.Println(size)
	}else{
		fmt.Errorf(string(err.Error()))
	}
}

func generateRandomString(strlen int) string {

	arr := make([]byte, strlen)
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	for x := 0; x < strlen; x++ {
		rand.Seed(time.Now().UnixNano())
		arr[x] = chars[rand.Intn(len(chars))]
	}
	fmt.Println("String which is generated is:", string(arr))
	return string(arr)
}
func main() {
	//fileIO()
	fileIOWithFlags()
}
