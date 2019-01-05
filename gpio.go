package gpio

import (
  "math"
  "os"
  "strconv"
  "time"
)

// gpio.GpioPin: provides support for opening and closing pins, and setting reading from pins.

type GpioPin struct {
  Number uint8
  IsWrite bool
  sysFile File
}

func (pin *GpioPin) Open(Pin uint8, Write bool) bool {
  // Check that the pin number is valid.
  if (Pin < 1 || Pin > 27) {
    return false
  }
  pin.Number = Pin
  // Inform the kernel that we wish to export the pin.
  file, error := os.Open("/sys/class/gpio/export")
  if error != nil {
    return false
  }
  file.WriteString(strconv.Itoa(Pin))
  file.Close()
  // Set the pin as input or output.
  file, error := os.Open("/sys/class/gpio/gpio" + strconv.Itoa(Pin) + "/direction")
  if error != nil {
    return false
  }
  if Write == true {
    file.Write("out")
  } else {
    file.Write("in")
  }
  file.Close()
  pin.IsWrite = Write
  // Open the pin for reading and writing.
  pin.sysFile, error := os.Open("/sys/class/gpio/gpio" + strconv.Itoa(Pin) + "/value")
  return true
}

func (pin *GpioPin) Write(Value bool) bool {
  // Make sure the pin was opened in the right mode.
  if pin.IsWrite == false {
    return false
  }
  // Write the desired value.
  if Value == true {
    pin.sysFile.WriteString("1")
  } else {
    pin.sysFile.WriteString("0")
  }
  return true
}

func (pin *GpioPin) Read() bool {
  // Read the value.
  var Value [1]byte
  result, error := pin.sysFile.Read(Value)
  // Interpret the value.
  if Value == "1" {
    return true
  } else {
    return false
  }
}

func (pin *GpioPin) Close() bool {
  // Close the value file.
  error := pin.sysFile.Close()
  if error != nil {
    return false
  }
  // Inform the kernel that we don't want the pin any more.
  file, error := os.Open("/sys/class/gpio/unexport")
  if error != nil {
    return false
  }
  file.WriteString(strconv.Itoa(pin.number))
  file.Close()
}

// gpio.GpioWriter: provides support for writing bytes to pins.

type GpioWriter struct {
  Pin GpioPin
  Baud uint
  baudController Ticker
}

func (writer *GpioWriter) Open(Number uint8, BitRate uint) bool {
  // Open the pin for writing.
  var pin GpioPin
  if pin.Open(Number, true) == true {
    writer.Pin = pin
  } else {
    return false
  }
  // Set the baud value.
  writer.Baud = BitRate
  // Set the pin to the default voltage.
  writer.Pin(Write(true))
  return true
}

func (writer *GpioWriter) WriteByte(Data uint8) {
  // Set up the ticker to tell us when to send a new bit.
  writer.baudController = time.NewTicker(time.Second / float32(Baud))
  // Set up a bitmask to extract bits from the byte.
  var bitMask byte = 1
  // Keep track of where we are in the byte.
  for bitNumber := -1; bitNumber <= 8; bitNumber++ {
    if bitNumber == -1 {
      // If we haven't started, signal the start of the byte.
      writer.Pin.Write(false)
    } else if bitNumber == 8 {
      // If we've finished, signal the end of the byte.
      writer.Pin.Write(true)
    } else {
      // Wait for the next tick.
      <- writer.baudController.C
      // Extract the bit and send it.
      writer.Pin.Write(Data & bitMask)
      // Advance the bitmask to the next bit.
      bitMask = bitMask << 1
    }
  }
  // Stop the ticker.
  writer.baudController.Stop()
  return
}

func (writer *GpioWriter) WriteBytes(Data []uint8) {
  // Construct a list of bits to send, including signal bits.
  var rawBits []bool
  var bitMask byte
  for i := 0; i < len(Data); i++ {
    // Signal bit: start byte
    rawBits = append(rawBits, false)
    // Data bits.
    bitMask = 1
    for j := 0; j < 8; j++ {
      rawBits = append(rawBits, Data[i] & bitMask)
      bitMask = bitMask << 1
    }
    // Signal bit: end byte
    rawBits = append(rawBits, true)
  }
  // Set up the ticker to tell us when to send a new bit.
  writer.baudController = time.NewTicker(time.Second / float32(Baud))
  // Send the bits.
  var bitsSent int64
  for bitNumber := 0; bitNumber < len(rawBits); bitNumber++ {
    // Wait for the next tick.
    <- writer.baudController.C
    // Send the bit.
    if writer.Pin.Write(rawBits[bitNumber]) == true {
      bitsSent++
    }
  }
  // Stop the ticker.
  writer.baudController.Stop()
  // Return the number of bytes sent.
  return bitsSent / 10
}

func (writer *GpioWriter) Close() bool {
  return writer.Pin.close()
}

// gpio.GpioReader: provides support for reading bytes from pins.

type GpioReader struct {
  Pin GpioPin
  Baud uint
  baudController Ticker
}

func (reader *GpioReader) Open(Number uint8, BitRate uint) bool {
  // Open the pin for reading.
  var pin GpioPin
  if pin.Open(Number, false) == true {
    reader.Pin = pin
  } else {
    return false
  }
  // Set the baud value.
  reader.Baud = BitRate
  return true
}

func (reader *GpioReader) ReadByte() byte {
  // Set up the ticker to tell us when to check for bits.
  reader.baudController = time.NewTicker(time.Second / float32(Baud))
  // Wait for the start signal.
  currentValue := true
  for ; currentValue == true {
    currentValue <- reader.Pin.Read()
  }
  // Read values.
  var output byte
  for i := 1; i <= 256; i = math.Pow(i, 2) {
    currentValue <- reader.Pin.Read()
    if currentValue == true {
      output += i
    }
  }
  return output
}

func (reader *GpioReader) Close() bool {
  return reader.Pin.Close()
}
