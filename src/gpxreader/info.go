/**
	GPX reader
*/

package gpxreader

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type Gpx struct {
	Creator string `xml:"creator,attr"`
	Time string `xml:"metadata>time"`
	Title string `xml:"trk>name"`
	TrackPoints []TrackPoint `xml:"trk>trkseg>trkpt"`
}

// string representation of the Gpx file
func (g Gpx) String() string {

  msg := `
GPX Reader v0.1.0
Creator:    %s
Time:       %s
Title:      %s
Max HR:     %v
Avg HR:     %v
`
	return fmt.Sprintf(msg, g.Creator, g.Time, g.Title, g.MaxHeartRate(), g.AverageHeartRate())
}

// represents each trackpoint
type TrackPoint struct {
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
	Elevation float32 `xml:"ele"`
	Time string `xml:"time"`
	Temperature int `xml:"extensions>TrackPointExtension>atemp"`
	HeartRate int `xml:"extensions>TrackPointExtension>hr"`
	Cadence int `xml:"extensions>TrackPointExtension>cad"`
}

// string representation of the Gpx file
func (t TrackPoint) String() string {
	return fmt.Sprintf("Lat: %v, Lon: %v", t.Lat, t.Lon)
}

// get the maximum heart rate
func (g Gpx) MaxHeartRate() (hr int) {
	hr = 0
	for _, pt := range g.TrackPoints {
		if pt.HeartRate > hr {
			hr = pt.HeartRate
		}
	}
	return
}

// get the average heart rate
func (g Gpx) AverageHeartRate() (hr int) {

  total := 0
  for _, pt := range g.TrackPoints {
    total += pt.HeartRate
  }

  hr = total / len(g.TrackPoints)
	return
}

func GpxFile(fileName string) (g Gpx) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
    fmt.Printf("FileName: %s", fileName)
		panic(err)
	}
	xml.Unmarshal(data, &g)

	return
}
