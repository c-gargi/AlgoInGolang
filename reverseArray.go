// Write a function to print an array in reverse order

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func reverseArray(a []int32) []int32 {

 for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {  
        a[i], a[j] = a[j], a[i]
    }

    return a
}

func main() {


    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    writer := bufio.NewWriter(os.Stdout)

    fmt.Print("Enter length of array : ")
    arrCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    fmt.Print("Enter the numbers to array : ")
    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(arrCount); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    res := reverseArray(arr)

    for i, resItem := range res {
        fmt.Fprintf(writer, "%d", resItem)

        if i != len(res) - 1 {
            fmt.Fprintf(writer, " ")
        }
    }

    fmt.Fprintf(writer, "\n")

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
