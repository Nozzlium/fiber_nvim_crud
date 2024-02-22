package main

import "github.com/nozzlium/fiber_nvim_crud/app"

func main() {
	err := app.GetApp().Listen("localhost:2637")
	if err != nil {
		panic(err)
	}
}
