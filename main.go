package main

import (
	"log"

	"github.com/ZetaMero/twittor/bd"
	"github.com/ZetaMero/twittor/handlers"

)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
