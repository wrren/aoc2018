package util

import (
	"fmt"
	"os"
)

func Fatal(err error) {
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
}
