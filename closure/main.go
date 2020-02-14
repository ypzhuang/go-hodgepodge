// closure an example to using closure
package main

import "fmt"

func main() {
	numGenerator := generator()
	// fmt.Println(numGenerator())
	// fmt.Println(numGenerator())
	// fmt.Println(numGenerator())
	// fmt.Println(numGenerator())
	// fmt.Println(numGenerator())
	for i := 0; i < 5; i++ {
		fmt.Println(numGenerator())
	}
}

func generator() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}
