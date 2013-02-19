/**
	GPX reader
*/

package gpxreader

import (
  "io/ioutil"
  "encoding/xml"
  "fmt"
  "time"
)

type Gpx struct {
  Creator     string `xml:"creator,attr"`
  Time        time.Time `xml:"metadata>time"`
  Title       string `xml:"trk>name"`
  TrackPoints []TrackPoint `xml:"trk>trkseg>trkpt"`
}

// string representation of the Gpx file
func (g Gpx) String() string {

  msg := `
GPX Reader v0.1.0
Creator:          %s
Time:             %s
Title:            %s
Max HR:           %v
Avg HR:           %v
Total Elevation:  %v
Cadence (Max):    %v
Cadence (Avg):    %v
`
  cadAvg, cadMax := g.Cadence()
  hrMax, hrAvg := g.HeartRate()
  return fmt.Sprintf(msg, g.Creator, g.Time, g.Title, hrMax,
    hrAvg, g.ElevationGain(), cadMax, cadAvg)
}

// represents each trackpoint
type TrackPoint struct {
  Lat         float64 `xml:"lat,attr"`
  Lon         float64 `xml:"lon,attr"`
  Elevation   float32 `xml:"ele"`
  Time        time.Time `xml:"time"`
  Temperature int `xml:"extensions>TrackPointExtension>atemp"`
  HeartRate   int `xml:"extensions>TrackPointExtension>hr"`
  Cadence     int `xml:"extensions>TrackPointExtension>cad"`
}

// string representation of the Gpx file
func (t TrackPoint) String() string {
  return fmt.Sprintf("Lat: %v, Lon: %v", t.Lat, t.Lon)
}

// get the maximum heart rate
func (g Gpx) HeartRate() (max, avg int) {
  max = 0
  total := 0
  for _, pt := range g.TrackPoints {
    total += pt.HeartRate
    if pt.HeartRate > max {
      max = pt.HeartRate
    }
  }
  avg = total / len(g.TrackPoints)
  return
}

func (g Gpx) ElevationGain() (gain float32) {

  el := g.TrackPoints[0].Elevation

  for i := 1; i < len(g.TrackPoints); i++ {
    ev := g.TrackPoints[i].Elevation
    if ev > el {
      gain += ev - el
    }
    el = ev
  }

  return
}

// Get the cadence stats.
// returns the average and max values
func (g Gpx) Cadence() (avg, max int) {

  count, total := 0, 0
  for _, pt := range g.TrackPoints {
    if pt.Cadence > 0 {
      count += 1
      total += pt.Cadence
    }

    if pt.Cadence > max {
      max = pt.Cadence
    }
  }

  avg = total / count

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
