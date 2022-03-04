package main

import "fmt"

func main() {
	fmt.Print("CHR\tDEC\tHEX\tOCT\n\n")
	for i := 33; i < 91; i++ {
		fmt.Printf("%c\t%d\t%x\t%O\n", i, i, i, i)
	}

}
