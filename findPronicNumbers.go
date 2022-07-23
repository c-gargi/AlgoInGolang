
// Write a function solution that, given two integers A and B returns the number of integers from the range [A….B](ends are included) which can be expressed as the product os two consecutive integers, that is X*(X+1), for some integer X.
// Examples:
// 1. Given A = 6 and B = 20, the function should return 3. These integers are 6 = 2*3, 12 = 3*4 and 20 = 4*5
// 2. Given A = 21 and B = 29, the function should return 0. There are no integers of the form X*(X+1) in this interval.
// Write an efficient algorithm for the following assumptions:
// * A and B are integers within the range [1.. 1,000,000,000];
// * A<=B.

package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
	"strconv"
	"strings"
	"math"
)

func checkPronic (number int64) bool {

 	num := int64(math.Floor(math.Sqrt(float64(number))))
    if (num * (num + 1) == number)  {
        return true 
    } else {
        return false 
    }

}

func main (){

	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	fmt.Print("Enter starting number : ")

	startNum, err := strconv.ParseInt(readLine(reader), 10, 64) /// 10 base, 64 bit size

	checkError(err)

	fmt.Print("Enter starting number : ")

	endNum, err := strconv.ParseInt(readLine(reader), 10, 64) /// 10 base, 64 bit size

	checkError(err)

	var pronicCount int 

	for i := startNum; i <= endNum; i++ {

		if isPronic := checkPronic(i); isPronic == true {
			pronicCount++
		}
	}

	fmt.Fprintf(writer, "%d\n", pronicCount)

  	writer.Flush()

}


func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }
    return strings.TrimRight(string(str), "\r\n")

}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}