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
  "gpx"
  "flag"
  "fmt"
)

func main() {

  var file = flag.String("filename", "", "Path to GPX file")

  flag.Parse()
  _, err := os.Stat(*file)
  if err != nil {
    flag.Usage()
    os.Exit(1)
  }

  gpx := gpx.GpxFile(*file)
  fmt.Print(gpx)

}

