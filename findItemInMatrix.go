// Write a function solution that finds needle in haystack 

package main

import (
	"fmt"
)


func searchNeedle (matrix [][]int, matrixSize int, needle int) (int, int) {

	var j int =  matrixSize -1
	for i:= 0; i<=matrixSize; i++{
		if matrix[i][j] == needle {
			return i, j 
		} else {
			
			j --
		}
	}
 	return 0,0
}


func main () {
 	matrix := [][]int{
 		{ 10, 20, 30, 40 }, 
        { 15, 25, 35, 45 }, 
       	{ 27, 29, 37, 48 }, 
      	{ 32, 33, 39, 50 }, 
 	}

 	matrixSize := 4
 	needle := 29

 	i, j := searchNeedle(matrix, matrixSize, needle)
 	fmt.Println("Found At ", i, j)

}