package main

import (
	"fmt"
)

func main()  {
	var fuct int = 1
	var n int
	fmt.Scan(&n)
	
	for i := 1; i <= n; i++ {
		fuct = fuct * i
	}

	fmt.Println(fuct)
}