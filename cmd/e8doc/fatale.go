package main

import (
	"log"
)

func fatale(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
