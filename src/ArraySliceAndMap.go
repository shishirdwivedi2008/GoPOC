// ArraySliceAndMap
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Function whose return type is array.
//Function should return with type and memory size as well.
func arrayFunction() (arr1 [100]int64) {
	var arr [100]int64
	for x := 0; x < 100; x++ {
		rand.Seed(time.Now().UnixNano())
		number := rand.Int63()
		fmt.Println("Random number generated", number)
		arr[x] = number
		time.Sleep(10 * time.Nanosecond)
	}
	return arr
}

func multiDimenisonalArray() {
	var newArray [100][100]int64
	for x := 0; x < len(newArray); x++ {
		for j := 0; j < len(newArray[0]); j++ {
			rand.Seed(time.Now().UnixNano())
			number := rand.Int63()
			newArray[x][j] = number
			time.Sleep(10 * time.Nanosecond)
		}
	}
	fmt.Print(newArray)
}

//Slices

func sliceExample() {
	lenOfString := 10
	mySlice := make([]string, 0)
	mySlice = append(mySlice, "shishir")
	fmt.Println(mySlice, "Lenght of slice", len(mySlice))
	mySlice = append(mySlice, "Dwivedi")
	for x := 0; x < 100; x++ {
		mySlice = append(mySlice, generateRandomString(lenOfString))
	}
	fmt.Print(mySlice)
	x := mySlice[:10]
	fmt.Println(x)
	x = mySlice[:]
	fmt.Println(x)
	x = mySlice[20:]
	fmt.Println(x)

}
func generateRandomString(strlen int) string {
	arr := make([]byte, strlen)
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	for x := 0; x < strlen; x++ {
		rand.Seed(time.Now().UnixNano())
		arr[x] = chars[rand.Intn(len(chars))]
		time.Sleep(10 * time.Nanosecond)
	}
	return string(arr)
}
func mapExample() {
	mapList := make(map[string]int)
	//adding random string as key and random int as value
	for x := 0; x < 100; x++ {
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(x)
		mapList[generateRandomString(10)] = number
		time.Sleep(10 * time.Nanosecond)
	}
	//Printing Map using Range Iterator.
	fmt.Println(mapList)

}

func main() {
	fmt.Println("Array following number:", arrayFunction())
	//multiDimenisonalArray()
	//sliceExample()
	mapExample()
}
