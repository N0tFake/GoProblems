package main

import (
	"cli/app"
	"log"
	"os"
)

func main() {
	aplication := app.Gerar()
	if err := aplication.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
