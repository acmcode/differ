package main

import (
	"fmt"
	"os"

	"github.com/acmcode/differ/tool"
)

// The entrance of the app
func main() {
	if len(os.Args) == 5 {
		userOut := os.Args[1]
		dataOut := os.Args[2]
		diffIgnoreHead := os.Args[3]
		strictMode := os.Args[4]
		isSame, err := tool.DiffOut(userOut, dataOut, diffIgnoreHead == "true", strictMode == "true")
		if err != nil {
			fmt.Println(err)
		} else {
			if isSame {
				fmt.Printf("'%s' is the same as '%s'\n", userOut, dataOut)
			} else {
				fmt.Printf("'%s' is not the same as '%s'\n", userOut, dataOut)
			}
		}
	} else {
		fmt.Println("Usage: ./acmdiffer user.out data.out false false")
	}
}
