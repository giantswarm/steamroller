package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path"
)

// SiZE defines the size of the dummy data written to file.
const SIZE = 4 * 1024 * 1024

// data holds the dummy data that is written to file.
var data []byte

var fName = flag.String("filename", "dummy-*.data", "filename template for dummy files")

// init initializes dummy data.
func init() {
	data := make([]byte, SIZE)
	for i := 0; i < SIZE; i++ {
		data[i] = 'A'
	}
}

func main() {
	flag.Parse()

	dir := path.Dir(*fName)
	file := path.Base(*fName)
	for {
		f, err := ioutil.TempFile(dir, file)
		if err != nil {
			panic(err)
		}

		for bs := 0; bs < len(data); {
			bsWritten, err := f.Write(data[bs:])
			if err != nil {
				panic(err)
			}

			bs = +bsWritten
		}

		err = f.Sync()
		if err != nil {
			panic(err)
		}

		f.Close()
		err = os.Remove(f.Name())
		if err != nil {
			panic(err)
		}
	}
}
