package main

import "fmt"

func main() {

	fmt.Println("Square", "Cube")
	fmt.Println(PowerSeries(4))
	fmt.Println(PowerSeries(5))

}

func PowerSeries(a int) (int, int, error) {
	return a*a, a*a*a, nil
}
