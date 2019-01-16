package main

import (
  "io/ioutil"
  "math"
  "strconv"
  "fmt"
)

func PrintPbmCFromFile(filename string) {
  rawbytes, error := ioutil.ReadFile(filename)
  if error != nil {
    panic(error)
  }
  if !(rawbytes[0] == 80 && rawbytes[1] == 52) {
    fmt.Println("ERROR: " + filename + " is not a valid uncompressed PBM file.")
    return
  }
  var widthString, heightString string
  var width, height int64
  var doneWidth, doneHeight bool
  for i := 3; doneHeight == false; i++ {
    if rawbytes[i] == 32 {
      doneWidth = true
    } else if rawbytes[i] == 10 {
      doneHeight = true
    } else {
      if doneWidth == false {
        widthString += strconv.Itoa(int(rawbytes[i]) - 48)
      } else {
        heightString += strconv.Itoa(int(rawbytes[i]) - 48)
      }
    }
  }
  width, _ = strconv.ParseInt(widthString, 10, 32)
  height, _ = strconv.ParseInt(heightString, 10, 32)
  fmt.Println("Reported: " + strconv.FormatInt(int64(width), 10) + " × " + strconv.FormatInt(int64(height), 10) + "px.")

  var dataBytes []byte
  var lfCount byte
  for i := 0; i < len(rawbytes); i++ {
    if lfCount < 2 {
      if rawbytes[i] == 10 {
        lfCount++
      }
    } else {
      dataBytes = append(dataBytes, rawbytes[i])
    }
  }

  var xL, xH, yL, yH byte
  xH = byte(math.Floor(float64(width) / 8 / 256))
  yH = byte(math.Floor(float64(height) / 256))
  xL = byte(float64(width - int64(xH)) / 8)
  yL = byte(height) - yL
  fmt.Println([]byte{xL, xH, yL, yH})
  fmt.Println(len(dataBytes))
  var outputBytes []byte = []byte{29, 118, 48, 0}
  outputBytes = append(outputBytes, xL, xH, yL, yH)
  outputBytes = append(outputBytes, dataBytes...)
  fmt.Println(outputBytes)
  SendBytes(outputBytes)
}

func main() {
  //SendCommand(CommandTestPage, []byte{})
  PrintPbmCFromFile("/home/cjs/latex.pbm")
  return
}
