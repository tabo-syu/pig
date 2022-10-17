package imager

import (
	_ "embed"
	"errors"
	"image"
	"image/color"

	"github.com/golang/freetype/truetype"
	"github.com/tabo-syu/pig/util"
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

func newTextDrawer(img *image.RGBA, ttf *truetype.Font, color image.Image, text string) *font.Drawer {
	size := img.Rect.Size()
	w, h := size.X, size.Y

	var (
		drawer    font.Drawer
		fontsize  float64 = float64(h) * 0.75
		textWidth int
	)
	for {
		drawer.Face = truetype.NewFace(ttf, &truetype.Options{Size: fontsize})
		textWidth = drawer.MeasureString(text).Ceil()
		if w/2 > textWidth {
			break
		}
		fontsize--
	}

	drawer.Dst = img
	drawer.Src = color
	drawer.Dot = fixed.Point26_6{
		X: fixed.I(w-textWidth) / 2,
		Y: fixed.I(h+int(fontsize/2)) / 2,
	}

	return &drawer
}

// Generate the placeholder image
func Generate(input Input) (*image.RGBA, error) {
	if err := validate(input); err != nil {
		return nil, err
	}

	ttf, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	r, g, b, a := util.ColorCodeToRGBA(input.Color)
	br, bg, bb := util.CalcColorFromBGColor(r, g, b)

	img := image.NewRGBA(image.Rect(0, 0, int(input.Width), int(input.Height)))
	fillColor(img, color.RGBA{r, g, b, a})
	drawer := newTextDrawer(img, ttf, image.NewUniform(color.RGBA{br, bg, bb, a}), input.Text)
	drawer.DrawString(input.Text)

	return img, nil
}
