package util

// Convert color code to RGBA
func ColorCodeToRGBA(code uint32) (r, g, b, a uint8) {
	r = uint8(code >> 24)
	g = uint8(code >> 16)
	b = uint8(code >> 8)
	a = uint8(code)

	return
}

// ITU-R Rec BT.601 Grayscale
func Grayscale(ir, ig, ib uint8) (r, g, b uint8) {
	gray := uint8(0.299*float32(ir) + 0.587*float32(ig) + 0.114*float32(ib))

	return gray, gray, gray
}

// Calculate text color from background color
func CalcColorFromBGColor(ir, ig, ib uint8) (r, g, b uint8) {
	gray, _, _ := Grayscale(ir, ig, ib)
	if gray > 186 {
		return 0, 0, 0
	}

	return 0xff, 0xff, 0xff
}
