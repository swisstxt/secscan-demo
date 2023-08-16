package main

import (
	"bytes"
	"math/rand"
	"golang.org/x/net/html"
)

func main() {
	source := bytes.NewBufferString(`<html><body><p>&</p></body></html>`)
	parsed, err := html.Parse(source)
	if err != nil {
		panic(err)
	}
	buffer := &bytes.Buffer{}
	html.Render(buffer, parsed)
	rand.Int()
}
