package main

import (
  "io/ioutil"
)

func SendString(Data string) {
  ioutil.WriteFile("/dev/ttyS0", []byte(Data), 660)
}

func SendBytes(Data []byte) {
  ioutil.WriteFile("/dev/ttyS0", Data, 660)
}

func SendCommand(command Command, data []byte) {
  SendBytes(command.ControlSequence)
  if len(data) > 0 {
    SendBytes(data)
  }
  if command.PostSequence != nil {
    SendBytes(command.PostSequence)
  }
}
