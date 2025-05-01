package main

import (
	"fmt"
	"image"
	"log"
	"math"

	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
)

func main() {
	fmt.Println("🎨 Color Picker")
	fmt.Println("Hover your mouse over a pixel and press Enter to pick a colour")

	fmt.Scanln()

	getColour()

	fmt.Println("📋 Hex value copied to clipboard.")

	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}

func getColour() {
	// Get mouse coordinates
	x, y := robotgo.Location()
	fmt.Printf("Mouse position: (%d, %d)\n", x, y)

	// Get display that contains the point
	displayIndex := -1
	for i := 0; i < screenshot.NumActiveDisplays(); i++ {
		bounds := screenshot.GetDisplayBounds(i)
		if pointInBounds(x, y, bounds) {
			displayIndex = i
			break
		}
	}
	if displayIndex == -1 {
		log.Fatal("Could not find display containing the cursor")
	}

	// Take screenshot of the display
	img, err := screenshot.CaptureRect(screenshot.GetDisplayBounds(displayIndex))
	if err != nil {
		log.Fatalf("Failed to capture screenshot: %v", err)
	}

	// Get color at (x, y) relative to the screen
	pixelColor := img.At(x-screenshot.GetDisplayBounds(displayIndex).Min.X, y-screenshot.GetDisplayBounds(displayIndex).Min.Y)
	r, g, b, _ := pixelColor.RGBA()
	r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

	// Convert to hex and HSL
	hex := fmt.Sprintf("#%02X%02X%02X", r8, g8, b8)
	h, s, l := rgbToHsl(r8, g8, b8)

	// Print values
	fmt.Println("Picked Color:")
	fmt.Printf("  RGB : %d, %d, %d\n", r8, g8, b8)
	fmt.Printf("  HEX : %s\n", hex)
	fmt.Printf("  HSL : %.0f, %.0f, %.0f\n", h, s, l)

	fmt.Println()
	fmt.Printf("  \033[38;2;%d;%d;%dm%s\033[0m\n", r8, g8, b8, "■■■■■■■■■■■■■■■■■■■■■■■■■■")

	// Copy hex to clipboard
	err = clipboard.WriteAll(hex)
	if err != nil {
		log.Fatalf("Failed to copy to clipboard: %v", err)
	}
}

func pointInBounds(x, y int, bounds image.Rectangle) bool {
	return x >= bounds.Min.X && x < bounds.Max.X && y >= bounds.Min.Y && y < bounds.Max.Y
}

func rgbToHsl(r, g, b uint8) (h, s, l float64) {
	rf, gf, bf := float64(r)/255, float64(g)/255, float64(b)/255
	max := math.Max(rf, math.Max(gf, bf))
	min := math.Min(rf, math.Min(gf, bf))
	l = (max + min) / 2

	if max == min {
		h, s = 0, 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case rf:
			h = (gf - bf) / d
			if gf < bf {
				h += 6
			}
		case gf:
			h = (bf-rf)/d + 2
		case bf:
			h = (rf-gf)/d + 4
		}
		h /= 6
	}

	return h * 360, s * 100, l * 100
}
