

package gpxreader


import (
	"io"
	"os"
	"encoding/xml"
)


// metadata type
type meta struct {
	creator string
	time string
	name string
}


// read a files contents
func readfile(file string) {

	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	var m meta
	xml.Unmarshal(f, &m)

	return
}


// get meta information
func Meta(dataFile string) (m meta) {

	return
}
