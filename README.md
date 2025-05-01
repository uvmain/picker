# 🎨 Color Picker

A simple Go-based application that allows you to pick a color from your screen by hovering your mouse over a pixel and pressing Enter. The application captures the color under the cursor, displays its RGB, HEX, and HSL values, and copies the HEX value to your clipboard.

## Features

- Capture the color of any pixel on your screen.
- Display the color in RGB, HEX, and HSL formats.
- Copy the HEX value of the color to your clipboard.
- Visualize the picked color in the console.

## Dev Prerequisites

- Go installed on your system.
- The following Go packages:
  - `github.com/atotto/clipboard`
  - `github.com/go-vgo/robotgo`
  - `github.com/kbinani/screenshot`

You can install these dependencies using:

```bash
go get github.com/atotto/clipboard
go get github.com/go-vgo/robotgo
go get github.com/kbinani/screenshot
```

## Usage

1. Clone this repository.
2. Run the application:

   ```bash
   go run .
   ```

3. Follow the instructions in the terminal:
   - Hover your mouse over the pixel you want to pick.
   - Press Enter to capture the color.
4. The application will display the color's RGB, HEX, and HSL values in the terminal.
5. The HEX value will be automatically copied to your clipboard.
6. Press Enter to exit the application.

## Example Output

```
🎨 Color Picker
Hover your mouse over a pixel and press Enter to pick a colour

Mouse position: (500, 300)
Picked Color:
  RGB : 255, 0, 0
  HEX : #FF0000
  HSL : 0, 100, 50

  ■■■■■■■■■■■■■■■■■■■■■■■■■

📋 Hex value copied to clipboard.

Press Enter to exit...
```

## Building

1. go build -ldflags "-s -w"
2. (optionally to reduce binary size, if you have upx installed) upx .\picker.exe

## License

This project is licensed under the MIT License. See the LICENSE file for details.
