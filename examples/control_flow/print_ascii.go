package main

import "fmt"

func main() {
	fmt.Print("CHR\tUNI\tDEC\tHEX\tOCT\tBIN\n\n")
	for i := 33; i < 122; i++ { //https://en.wikipedia.org/wiki/ASCII
		fmt.Printf("%c\t%#U\t%v\t%#x\t0x%o\t%b\n", i, i, i, i, i, i)
	}
}
