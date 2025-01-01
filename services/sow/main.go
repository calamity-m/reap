package main

import (
	"fmt"
	"os"

	"github.com/calamity-m/reap/services/sow/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("failed to run sow")
		os.Exit(1)
	}
}
