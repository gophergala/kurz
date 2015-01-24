package main

import (
	"github.com/FGM/kurz/storage"
	"github.com/davecgh/go-spew/spew"
	"log"
)

func main() {
	err := storage.Service.Open("someuser:somepass@tcp(localhost:3306)/go_kurz")
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Service.Close()

	spew.Dump(storage.StrategyStatistics(storage.Service))
}
