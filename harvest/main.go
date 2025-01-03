package main

import (
	"fmt"
	"os"

	"github.com/calamity-m/reap/harvest/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("failed to run harvest")
		os.Exit(1)
	}
}
