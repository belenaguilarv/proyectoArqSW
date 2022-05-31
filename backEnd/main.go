package main

import (
	"github.com/belenaguilarv/proyectoArqSW/backEnd/app"
	"github.com/belenaguilarv/proyectoArqSW/backEnd/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
