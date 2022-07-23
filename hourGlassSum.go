// Given a 6 x 6 2D Array define an hourglass to be a subset of values. 
// Calculate the hourglass sum for every hourglass in the array , then print the maximum hourglass sum.

package main
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

func hourglassSum(arr [][]int32) int32 {

var maxSum int32
maxSum = math.MinInt32

  for  row := 0; row <= 3; row++ {
               for col := 0; col <= 3; col++ {
                    var sum int32
          
                    sum += arr[row][col]
                    sum += arr[row][col + 1]
                    sum += arr[row][col + 2]
                    sum += arr[row + 1][col + 1]
                    sum += arr[row + 2][col]
                    sum += arr[row + 2][col + 1]
                    sum += arr[row + 2][col + 2]
          
                    if sum > maxSum {
                        maxSum = sum
                    }
               } 
          }

    return maxSum
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    writer := bufio.NewWriter(os.Stdout)
    var arr [][]int32
    fmt.Print("Enter the matrix : \n")

    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := hourglassSum(arr)

    fmt.Fprintf(writer, "%d\n", result)

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