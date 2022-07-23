// Write a function solution that given an integer N, returns the smallest number with the same number of digits. 
// N is between 1 and one billion.
// For example, given N = 125, the function should return 100. 
// Given N=10, the function should return 10. Given N=1, the function should return 0.


package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func digitsCount(number int32) int {
 count := 0
 for number != 0 {
  number /= 10
  count += 1
 }
 return count

}

func nextSmallest (number int32, digitLength int) int32 {

if number == 1 {
 return 0
}
var res int32

for i := number; i > 0; i-- {
       
    currentDigitLength := digitsCount(i)
	  if currentDigitLength == digitLength && i <= number {	
	    res = i	    
	  }	
}

return res
}


func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  writer := bufio.NewWriter(os.Stdout)

  fmt.Print("Enter number : ")

  inNum, err := strconv.ParseInt(readLine(reader), 10, 64)

  checkError(err)

  digitLength := digitsCount(int32(inNum))

  nextSmallestNum := nextSmallest(int32(inNum), digitLength)

  fmt.Fprintf(writer, "%d\n", nextSmallestNum)

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
