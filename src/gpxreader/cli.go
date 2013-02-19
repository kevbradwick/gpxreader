/**
 * Created with IntelliJ IDEA.
 * User: bradwk01
 * Date: 19/02/2013
 * Time: 12:52
 * To change this template use File | Settings | File Templates.
 */
package main

import (
  "os"
  "gpxreader"
  "fmt"
)

func main() {

  if len(os.Args) != 2 {
    fmt.Println("Please specify a file name")
    os.Exit(1)
  }

  gpx := gpxreader.GpxFile(os.Args[1])
  fmt.Println(gpx)
}

