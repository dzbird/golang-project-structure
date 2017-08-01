package main

import macaron "gopkg.in/macaron.v1"

func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello world2!"
	})
	m.Run()
}
