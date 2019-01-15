package main

type Command struct {
  Name string
  ControlSequence []byte
  DataLengthRange [2]uint64
  PostSequence []byte
  Description string
}

var CommandPrintAndLineFeed = Command{
  Name: "Print and line feed",
  ControlSequence: []byte{10},
  Description: "Prints the data in the print buffer and feeds one line, based on the current line spacing. This command sets the print position to the beginning of the line.",
}

var CommandPrintAndCarriageReturn = Command{
  Name: "Print and carriage return",
  ControlSequence: []byte{13},
  Description: "When automatic line feed is enabled, this command functions the same as Print and line feed; when automatic line feed is disabled, this command is ignored. Sets the print starting position to the beginning of the line.",
}

var CommandHorizontalTab = Command{
  Name: "Horizontal tab",
  ControlSequence: []byte{9},
  Description: "Moves the print position to the next horizontal tab position. This command is ignored unless the next horizontal tab position has been set. If the next horizontal tab position exceeds the printing area, the printer sets the printing position to [printing area width+1]. If this command is received when the printing position is at [printing area width+1], the printer executes print buffer-full printing of the current line and horizontal tab processing from the beginning of the next line.",
}

var CommandSetHorizontalTabPositions = Command{
  Name: "Set horizontal tab positions",
  ControlSequence: []byte{27, 68},
  DataLengthRange: [2]uint64{1, 32},
  PostSequence: []byte{0},
  Description: "Set horizontal tab positions. Data consists of a sequence of bytes representing the horizontal tab position, measured in character-width columns. This command deletes previous tab settings. Up to 32 tabs can be set. Transmit tabs in ascending order. Default tabs are every 8 characters.",
}

var CommandPrintAndFeedPaper = Command{
  Name: "Print and feed paper",
  ControlSequence: []byte{27, 74},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Prints the data in the print buffer and feeds the paper by 0.125 mm times the value of the byte following the control sequence. After printing is completed, this command sets the print starting position to the beginning of the line.",
}

var CommandPrintAndFeedNLines = Command{
  Name: "Name and feed n lines",
  ControlSequence: []byte{27, 100},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Print the data in the buffer and feed paper n lines, where n is the value of the byte following the control sequence. This command sets the print starting position to the beginning of the line. The maximum paper feed amount is 1016 mm (40 inches). If the paper feed amount (n×line spacing) of more than 1016 mm (40 inches) is specified, the printer feeds the paper only 1016 mm (40 inches).",
}

var CommandSetPeripheralDevice = Command{
  Name: "Set peripheral device",
  ControlSequence: []byte{27, 61},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Sets the device online or offline for receiving print data. This is set in the first bit of the byte following the control sequence: 0 is offline, 1 is online.",
}

var CommandSelectDefaultLineSpacing = Command{
  Name: "Select default line spacing",
  ControlSequence: []byte{27, 50},
  Description: "Selects 3.75 mm (30×0.125 mm) line spacing.",
}

var CommandSetLineSpacing = Command{
  Name: "Set line spacing",
  ControlSequence: []byte{27, 51},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Sets the line spacing to [n×0.125 mm], where n is the value of the byte following the control sequence.",
}

var CommandSelectJustification = Command{
  Name: "Select justification",
  ControlSequence: []byte{27, 97},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Aligns all the data in one line to the specified position, specified in the byte following the control sequence. Values of 0 or 48 align left, 1 or 49 align centre, and 2 or 50 align right. The command is enabled only when processed at the beginning of the line in standard mode.",
}

var CommandSetLeftMargin = Command{
  Name: "Set left margin",
  ControlSequence: []byte{29, 76},
  DataLengthRange: [2]uint64{2, 2},
  Description: "Sets the left margin using the two bytes of data (labelled nL, nH) following the control sequence. The left margin is set to [(nL+nH×256) ×0.125 mm]. This command is effective only when processed at the beginning of the line in standard mode. If the setting exceeds the printable area, the maximum value of the printable area is used.",
}

var CommandSetLeftBlankCharNumbers = Command{
  Name: "Set left blank char numbers",
  ControlSequence: []byte{27, 66},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Default is 0. Value of byte may be 0 ≤ m ≤ 47.",
}

var CommandSetAbsolutePrintPositions = Command{
  Name: "Set absolute print positions",
  ControlSequence: []byte{27, 36},
  DataLengthRange: [2]uint64{2, 2},
  Description: "Set the distance from the beginning of the line to the position at which subsequent [characters?] are to be printed. The distance from the beginning of the line to the print position is [(nL+nH×256)×0.125 mm] (data labelled nL, nH).",
}

var CommandSelectPrintMode = Command{
  Name: "Select print mode",
  ControlSequence: []byte{27, 33},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Selects print mode(s) using data byte as follows. Bit 0: 0 = Character Font A (12×24); 1 = Character Font B (9×17). Bit 1: 1 = Colour inversion on. Bit 2: 1 = Upside-down printing. Bit 3: 1 = Emphasised mode. Bit 4: 1 = Double height mode. Bit 5: 1 = Double width mode selected. Bit 6: 1 = Deleteline mode on.",
}

var CommandSelectCharacterSize = Command{
  Name: "Select character size",
  ControlSequence: []byte{29, 33},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Selects the character height using bits 0 to 2 and selects the character width using bits 4 to 7, as follows: the first four bits select character height; the second four bits select character width. Interpreting these four bits as nibbles, they represent a one-based multiplier for the character size, up to a maximum of eight.",
}

var CommandInvertColours = Command{
  Name: "Turn white-black reverse printing mode",
  ControlSequence: []byte{29, 66},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Turns on or off white/black reverse printing mode. The first bit of the data byte sets the mode on or off. This mode has priority over underlining.",
}

var CommandClockwiseRotate = Command{
  Name: "Turn 90° clockwise rotation mode on/off",
  ControlSequence: []byte{27, 86},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Turns 90° clockwise rotation mode on/off. The least significant bit sets the mode on or off.",
}

var CommandDoubleStrikeMode = Command{
  Name: "Turn on/off double-strike mode",
  ControlSequence: []byte{27, 71},
  DataLengthRange: [2]uint64{1, 1},
  Description: "LSB sets the double-strike mode on and off. Output identical to emphasised mode.",
}

var CommandEmphasiseMode = Command{
  Name: "Turn emphasized mode on/off",
  ControlSequence: []byte{27, 69},
  DataLengthRange: [2]uint64{1, 1},
  Description: "LSB sets the emphasised mde on and off.",
}

var CommandTracking = Command{
  Name: "Set right-side character spacing",
  ControlSequence: []byte{27, 32},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Sets the right-side tracking to [n×0.125mm (n×0.0049”), where n is the data byte.",
}

var CommandDoubleWidthMode = Command{
  Name: "Select double width mode",
  ControlSequence: []byte{27, 14},
  DataLengthRange: [2]uint64{0, 1},
  Description: "Data byte is listed, but not documented.",
}

var CommandDisableDoubleWidthMode = Command{
  Name: "Disable double width mode",
  ControlSequence: []byte{27, 20},
  DataLengthRange: [2]uint64{0, 1},
  Description: "Data byte is listed, but not documented.",
}

var CommandUpsideDownMode = Command{
  Name: "Turns on/off upside-down printing mode",
  ControlSequence: []byte{27, 123},
  DataLengthRange: [2]uint64{1, 1},
  Description: "LSB of data bit represents whether lines should be printed upside down.",
}

var CommandUnderlineMode = Command{
  Name: "Turn underline mode on/off",
  ControlSequence: []byte{27, 45},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Data byte sets underline mode. 0, 48 = off; 1, 49 = 1 dot thick; 2, 50 = 2 dots thick.",
}

// User-defined characters and japanese stuff.

var CommandSelectCharacterSet = Command{
  Name: "Select an internal character set",
  ControlSequence: []byte{27, 82},
  DataLengthRange: [2]uint64{1, 1},
  Description: "0-based incremented maps are: 0 = USA; France; Germany; UK; Denmark I; Sweden; Italy; Spain I; Japan; Norway; Denmark II; Spain II; Latin America; Korea; Slovenia/Croatia; 15 = China.",
}

var CommandSelectCharacterCodeTable = Command{
  Name: "Select character code table",
  ControlSequence: []byte{27, 116},
  DataLengthRange: [2]uint64{1, 1},
  Description: "Refer to documentation.",
}

// Something chinese.






var CommandPrintRaster = Command{
  Name: "Print raster data (internal alias)",
  ControlSequence: []byte{29, 118, 48, 48},
  DataLengthRange: [2]uint64{5, 4294967296},
  Description: "Prints raster data at 203.2 dpi. MSB at left; LSB at right. Data bytes are xL, xH, yL, yH, then pixel bytes. Image size specified by x = xL+xH×256, y = yL+yH×256.",
}

var CommandInitialise = Command{
  Name: "Initialise the printer",
  ControlSequence: []byte{27, 64},
  Description: "Initialises the printer. Print buffer is cleared, parameters reset to defaults, user-defined characters cleared, DIP switch settings not checked, receive buffer not cleared.",
}

var CommandStatus = Command{
  Name: "Transmit paper sensor status",
  ControlSequence: []byte{27, 118},
  Description: "Returns one byte. Bit 0: represents online/offline; Bit 2: represents paper status; Bit 3: Whether voltage exceeds 9.5V; Bit 6: Whether temperature exceeds 60 degC.",
}

var CommandTestPage = Command{
  Name: "Print test page",
  ControlSequence: []byte{18, 84},
  Description: "Prints a test page.",
}
