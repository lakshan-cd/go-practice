package main

import (
	"fmt"
)

//only will excute specif block not whole code and break will automaticlay provide
func switchBasic() {
	fmt.Print("Go runs on ")
	os:= "linux"
	switch os {
	case "darwin" :
		fmt.Println("macOS.")
	case "linux" :
		fmt.Println("Linux.")
	
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}