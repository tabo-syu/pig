package imager

import (
	_ "embed"
	"errors"
	"image"
	"image/color"

	"github.com/golang/freetype/truetype"
	"golang.org/x/exp/slices"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

type formats []string

const (
	Jpg string = "jpg"
	Png string = "png"
	Gif string = "gif"
)

var Formats formats = []string{Jpg, Png, Gif}

type Input struct {
	Width  uint16
	Height uint16
	Format string
	Text   string
	Color  uint32
}

var (
	ErrInvalidImageSize   = errors.New("width and height must be at least 1")
	ErrInvalidImageFormat = errors.New("specify a format in jpg, png, or gif")
)

// Validate input data
func validate(input Input) error {
	// validate image size
	if input.Width == 0 || input.Height == 0 {
		return ErrInvalidImageSize
	}

	// validate image format
	if !slices.Contains(Formats, input.Format) {
		return ErrInvalidImageFormat
	}

	// Round to the maximum 32-bit value
	if input.Color > 0xff_ff_ff_ff {
		input.Color = 0xff_ff_ff_ff
	}

	return nil
}

// Convert color code to RGBA
func colorCodeToRGBA(code uint32) (r, g, b, a uint8) {
	r = uint8(code >> 24)
	g = uint8(code >> 16)
	b = uint8(code >> 8)
	a = uint8(code)

	return r, g, b, a
}

// Fill the image with color
func fillColor(img *image.RGBA, c color.Color) {
	size := img.Rect.Size()
	width, height := size.X, size.Y
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, c)
		}
	}
}

// Insert text to image
func insertLabel(img *image.RGBA, color image.Image, text string) error {
	size := img.Rect.Size()
	w, h := size.X, size.Y

	ttf, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}
	var fontsize float64 = 100
	face := truetype.NewFace(ttf, &truetype.Options{Size: fontsize})

	drawer := &font.Drawer{
		Dst:  img,
		Src:  color,
		Face: face,
		Dot:  fixed.Point26_6{},
	}
	a := drawer.MeasureString(text)
	drawer.Dot = fixed.Point26_6{
		X: (fixed.I(w) - a) / 2,
		Y: fixed.I(h+(int(fontsize)/2)) / 2,
	}

	drawer.DrawString(text)

	return nil
}

// Generate the placeholder image
func Generate(input Input) (*image.RGBA, error) {
	if err := validate(input); err != nil {
		return nil, err
	}

	img := image.NewRGBA(
		image.Rect(0, 0, int(input.Width), int(input.Height)),
	)
	r, g, b, a := colorCodeToRGBA(input.Color)
	fillColor(img, color.RGBA{r, g, b, a})
	insertLabel(img, image.Black, input.Text)

	return img, nil
}
