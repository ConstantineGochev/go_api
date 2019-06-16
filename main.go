package main

import (
	app "github.com/ConstantineGochev/go_api/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
