/**
 * Created with IntelliJ IDEA.
 * User: kevin
 * Date: 16/02/2013
 * Time: 22:13
 * To change this template use File | Settings | File Templates.
 */
package main


import (
	"fmt"
	"gpxreader"
)


func main() {

	file := "./fixtures/data_001.gpx"
	gpx := gpxreader.GpxFile(file)
	fmt.Println(gpx)
	fmt.Println(gpx.MaxHeartRate())

}
