package main

import (
	"log"
	"log/slog"
)

func main() {

	sow := NewSowServer(*slog.Default(), "localhost:8099")

	err := sow.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
