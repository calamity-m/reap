package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	Uuid      uuid.UUID
	GoldStars int
	Value     string
	Expiry    time.Time
	Created   time.Time
}

func main() {
	fmt.Println("a")
}
