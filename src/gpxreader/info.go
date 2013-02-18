/**
	GPX reader
*/

package gpxreader

import (
	"io/ioutil"
	"encoding/xml"
)

type Gpx struct {
	Creator string `xml:"creator,attr"`
	Time string `xml:"metadata>time"`
	Title string `xml:"trk>name"`
	TrackPoints []TrackPoint `xml:"trk>trkseg>trkpt"`
}

// represents each trackpoint
type TrackPoint struct {

}

// string representation of the Gpx file
func (g Gpx) String() string {
	return g.Title
}

func GpxFile(fileName string) (g Gpx) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	xml.Unmarshal(data, &g)

	return
}
