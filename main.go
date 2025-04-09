package main

import (
	"github.com/Jubair-Ahmed-Khan/ToDoList/userView"
)

func main() {
	app := userView.NewCLI()
	app.Run()
}