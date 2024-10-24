package main

import (
	"cloudLogServer/modUtility"
	"fmt"
)

func main() {
	err := modUtility.Utility_Initialize()
	if err != nil {
		fmt.Println("utility init error")
	}

	err = systemInit()
	if err != nil {
		return
	}

}
