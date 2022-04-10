package main

// Importing fmt and path/filepath
import (
	"fmt"
	"os"
	"path/filepath"
)

// Calling main
func main() {
	joiningPathsWithCurrentWorkingDir()

}

func joiningPathsWithCurrentWorkingDir() {
	fmt.Println("Joining paths with the current working dir:")

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("\tWorking Dir:\t\t", pwd)

	abs, _ := filepath.Abs(pwd)
	fmt.Println("\tFull Path:\t\t", abs)

	f := filepath.Join(pwd, "web", "images", "gif")
	fmt.Println("\tSub Path off PWD:\t", f)
}
