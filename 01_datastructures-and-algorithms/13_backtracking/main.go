//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

func findCombWithSum(arr []int, curComb []int, curSum, target, curIndex int) int {
	var count = 0

	curSum += arr[curIndex]
	if curSum > target {
		return 0
	}
	curComb = append(curComb, arr[curIndex])
	if curSum == target {
		count += 1
		for _, v := range curComb {
			fmt.Printf("%d, ", v)
		}
		fmt.Println()
	}
	for i := curIndex; i< len(arr); i++ {

		count += findCombWithSum(arr, curComb, curSum, target, i)
	}
	return count

}

//findElementsWithSum  of k from arr of size
func findElementsWithSum(arr [2]int, combinations [4]int, size int, targetVal int, addValue int, l int, m int) int {

	var count int = 0

	if addValue > targetVal {
		return 0
	}

	if addValue == targetVal {
		count = count + 1
		var p int = 0
		for p = 0; p < m; p++ {

			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
		fmt.Println("m", m, "l", l)
		//time.Sleep(200 * time.Millisecond)
	}

	var i int
	for i = l; i < size; i++ {

		//fmt.Println(" m", m)
		combinations[m] = l

		count += findElementsWithSum(arr, combinations, size, targetVal, addValue+arr[i], l, m+1)
		l = l + 1
	}
	return count
}

// main method
func main() {

	////var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	//var arr = [2]int{1, 2}
	//
	//var addedSum int = 3
	//
	//var combinations [4]int
	//
	//count := findElementsWithSum(arr, combinations, 2, addedSum, 0, 0, 0)
	//
	////fmt.Println(check)
	//
	////var check2 bool = findElement(arr,9)
	//
	////fmt.Println(check2)
	//fmt.Println(count)
	var arr = []int{1, 2,3}
	var com = []int{}
	count := findCombWithSum(arr, com, 0, 3, 0)
	fmt.Println(count)
}
