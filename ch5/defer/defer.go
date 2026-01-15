package main

import (
	"io"
	"log"
	"os"
)

// Use defer keyword to run code after a function exits (successfully or unsuccessfully)
// One good use of named return params is that you can defer a function which modifies
// those return values.

// Example function cat
func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// ATP file is open so we should close it whenever the function ends
	defer f.Close()

	data := make([]byte, 2048)

	for {
		// read <count> bytes into <data> from <f>
		count, err := f.Read(data)
		// write data to stdout
		os.Stdout.Write(data[:count])

		// exit on error - if EOF, exit normally (i.e. no error)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

	}

}
