package imager

import (
	"errors"
	"image"
	"image/color"

	"golang.org/x/exp/slices"
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

	return nil
}

// Fill the image with color
func fillColor(img *image.RGBA, c color.Color) {
	size := img.Rect.Size()
	width := size.X
	height := size.Y
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, c)
		}
	}
}

// Insert text to image
func insertText(img *image.RGBA, text string) {

}

// Generate the placeholder image
func Generate(input Input) (*image.RGBA, error) {
	if err := validate(input); err != nil {
		return nil, err
	}

	img := image.NewRGBA(
		image.Rect(0, 0, int(input.Width), int(input.Height)),
	)
	fillColor(img, color.White)
	insertText(img, input.Text)

	return img, nil
}
