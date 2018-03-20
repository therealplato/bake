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
		log.Fatal(`github.com/therealplato/bake
make a go file containing filename's contents as a variable
usage: bake filename`)
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(filename + " not found")
	}
	outfile := filename + ".go"
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
	for {
		n, err = buf.ReadFrom(f)
		fmt.Printf("outer %v\n", n)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("failed reading g: " + err.Error())
		}

		m = m + n
		for {
			fmt.Printf("inner %v\n", m)
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
		buf.Reset()
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
