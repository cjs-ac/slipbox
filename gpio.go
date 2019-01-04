package gpio

import (
  "os"
  "strconv"
)

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
