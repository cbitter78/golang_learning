package main

import (
	"fmt"
)

func main() {
	// You should go read this: https://go.dev/blog/slices-intro

	// Array is the base type of slice. Most of the time we just use slice not array
	// REMEMBER An array’s size is fixed.
	// IMPORTANT: If you pass an array to a function then go will COPY every elements value
	//            to NEW memory and that will be passed to the function.
	//            WARNING!!!! This can cause Massive memory consumption and slowness
	//            if you're using a recursive function.  Pro-tip: Slices do not do this
	//            See: https://stackoverflow.com/questions/39993688/are-slices-passed-by-value
	a := [5]int{1, 3, 5, 8, 13}
	fmt.Printf("Created Array witn %v elements %v\n", len(a), a)
	for i, v := range a {
		fmt.Printf("\t a[%v]:%v\n", i, v)
	}

	// Lets play with Slice
	// A slice layers on top of array and has a len and capacity.  See https://go.dev/src/runtime/slice.go#L13
	// this means it just points to the array.  The capacity is the total size
	// of the underlying array and the len is how much of that array the slice is taking up or using.
	// as the slice grows (using append) it just moves the len pointer down the array which
	// is fast!  However it might hit the end of the underlying array cap(a) == len(a)
	// in this case the append method makes a new array to back the slice that is double
	// the size of the old one.  And it COPIES! yes copies!  all the data into that
	// new array, then throws away the old one.   Growing a slice beyond its capacity
	// `might` be expensive.
	s := []int{21, 34, 55, 89, 144, 233} // Notice we did not declair the size  :)
	fmt.Printf("Created Slice witn %v elements %v with the capacity of %v\n", len(s), s, cap(s))

	fmt.Println("Iterate over the slice using range to get the index and value")
	for i, v := range s {
		fmt.Printf("\t s[%v]:%v\n", i, v)
	}

	fmt.Println("Iterate over the slice using index")
	for i := 0; i <= len(s)-1; i++ {
		fmt.Printf("\t s[%v]:%v\n", i, s[i])
	}

	// Slicing a Slice.
	// IMPORTANT: Slicing does not copy the slice’s data.
	//            It creates a new slice value that points to the original array.
	//            This makes slice operations as efficient as manipulating array indices.
	//            Therefore, modifying the elements (not the slice itself) of a re-slice
	//            modifies the elements of the original slice:
	//
	// IMPORTANT: Sense a slice or a slice of a slice is a points to the underlying
	//            array the full array will be kept in memory until it is no longer referenced.
	//            Occasionally this can cause the program to hold all the data in memory when only a
	//            small piece of it is needed.  Use the copy command to avoide this.
	s1 := s[2:]  // Start at index 2 and go to the end
	s2 := s[1:3] // Start at index 1 then up to index 3 but not including it
	s3 := s[:3]  // Start at index 0 and go up to index 3 but not including it
	fmt.Println("Slicing the Slice:")
	fmt.Printf("\t1: the forth element to the end          %v\n", s1)
	fmt.Printf("\t2: the third second element to the forth %v\n", s2)
	fmt.Printf("\t3: the first three elements              %v\n", s3)
	fmt.Printf("\ts: the orginal slice                     %v\n\n", s)
	fmt.Println("Notice the 3rd element is 55 and each slice contains this element")
	fmt.Println("Update the original slice so the thrid element is differnt.")
	s[2] = 55555555
	fmt.Printf("\t1: the forth element to the end          %v\n", s1)
	fmt.Printf("\t2: the third second element to the forth %v\n", s2)
	fmt.Printf("\t3: the first three elements              %v\n", s3)
	fmt.Printf("\ts: the orginal slice                     %v\n\n", s)
	s[2] = 55 // Change it back.

	// You can convert an array to a Slice (or use it as the underlying array of a slice)
	// by slicing it.  In this case we slice array a and get the entire thing
	a_as_a_slice := a[:] // Start at The beginning : and go to the end.
	fmt.Println("Lets convert an Array to a Slice")
	printSlice("array (a) converted to a slice.  Now it has a capacity.", a_as_a_slice)

	fmt.Println("Lets append to a slice")
	// The append method takes a slice, then a bunch of elements to append to it and **Returns a new slice**
	s = append(s, 377, 610, 987)
	printSlice("s appended with 3 more values", s)

	// This is important because you cant just append 2 slices together like you can
	// with a python list.
	// For example this does NOT work!
	// s2 := []int {1,2,3}
	// s = append(s, s2)  // Uncomment this and you will see that it does not compile
	// this is because append takes a slice and then a LIST of elments e,e,e,e,e,e,e,e,e,e,e,e,e,e,e

	// So how do you do this if you ahve 2 slices?   you Unfurl the second slice.
	ss := []int{610, 987, 1597, 2584, 4181, 6765}
	s = append(s, ss...)
	// This is the same as s = append(s, 610, 987, 1597, 2584, 4181, 6765)
	// go just Unfurls the elements into the append function.
	// You should note that ss... will be sent to append as a slice and it
	// will point to the same underlying array that ss uses.
	printSlice("s appended to ss", s)

	// Why even make a slice?
	fmt.Println("Lets make some slices")
	var c []int
	printSlice("c empty", c)
	c = append(c, 1)
	printSlice("c grow by 1", c)
	c = append(c, 1)
	printSlice("c grow by 1 more", c)
	for i := 1; i < 49; i++ {
		c = append(c, 1)
		printSlice("c grow by 1 more", c)
	}
	println("At this point we have created and destroyed 49 arrays!  Not so awesome  :(\n")

	println("lets `make` a slice:")
	// Lets say we know we are going to process some data and we expect 100 elements
	// but if we get more than that we dont want to copy the array at 101 then again at
	// 102 then 103.   Maybe we can do better than that.  Say we want to ensure we have
	// a underylying array of 100 elements but we want to grow by 50 more at a time.
	// this is where make helps us.

	s_m := make([]int, 0, 100)
	printSlice("A `made` slice that has no elements but can grow to 100 without an array copy", s_m)
	for i := 0; i < 101; i++ {
		s_m = append(s_m, 1)

	}
	printSlice("s_m with 101 elements", s_m)
	println("Now we have a slice that has grown beyond its orignial size.")
	println("At this point the underlying array of size 100 has been destroyed")
	println("and replaced with a new one of size 200.   we can see this by slicing into it")
	printSlice("slice of s_m beyond where we put data:", s_m[80:150])
	// Rember the len of s_m is 101 and we are going beyond that

	println("You see where the 1's turn to 0's?  Thats the end of the data we were playing")
	println("with and the empty data from the underlying array.")

	// We use make so that we can change the memory allocation behind slice
	// in the case we are doing loop[n] or large data.  The go docs say this
	// covers 99% of the use cases.  However we are encuraged to implement
	// our own append function if we want it to act differntly.
	// TLDR; use make if your slice `might` grow beyond what you expect and just
	//       double the size or maybe trippe it.  ¯\_(ツ)_/¯
}

func printSlice(name string, slice []int) {
	// Passing a slice to a function is ok because it
	// DOES NOT copy the underlying array.  It does copy the
	// "header" or pointer, len, and cap values.  This is not
	// a huge deal.  In go EVERYTHING is passed by value not referance.  :)
	fmt.Printf("Slice %v:\n", name)
	fmt.Printf("\tlen:  %v\n", len(slice))
	fmt.Printf("\tcap:  %v\n", len(slice))
	fmt.Printf("\tdata: %v\n", slice)
	fmt.Println("")
}
