package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{},
	}

	app := application{
		config: cfg,
	}

	h := app.mount()
	err := app.run(h)
	if (err != nil){
		log.Println("Server has failed to start")
		os.Exit(1)
	}
}
