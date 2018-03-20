package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`github.com/therealplato/bake
make a go file containing filename's contents as a variable
usage: bake filename`)
		os.Exit(1)
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(filename + " not found")
	}
	outfile := "baked.go"
	g, err := os.Create(outfile)
	bake(f, g)
}

func bake(f io.Reader, g io.Writer) {
	var (
		m   int64
		n   int64
		b   byte
		err error
		buf bytes.Buffer
	)
	head(g)
	n, err = buf.ReadFrom(f)
	fmt.Printf("read %v bytes\n", n)
	for {
		b, err = buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("failed reading buf: " + err.Error())
		}
		_, err := g.Write([]byte(fmt.Sprintf("%#x, ", b)))
		if err != nil {
			log.Fatal("failed writing g: " + err.Error())
		}

		if m%20 == 0 {
			// 80 characters
			g.Write([]byte("\n"))
		}
	}
	tail(g)
}

func head(g io.Writer) {
	_, err := g.Write(
		[]byte(
			`package main

var baked = []byte{
`),
	)
	if err != nil {
		log.Fatal("failed writing g: " + err.Error())
	}
}

func tail(g io.Writer) {
	_, err := g.Write(
		[]byte("}\n"),
	)
	if err != nil {
		log.Fatal("failed writing g: " + err.Error())
	}
}
