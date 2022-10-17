/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/spf13/cobra"
	"github.com/tabo-syu/placeholder-image-generator/imager"
)

type Writers struct {
	Out io.Writer
	Err io.Writer
}

type RootFlags struct {
	Width  uint16
	Height uint16
	Format string
	Text   string
	Color  uint32
}

func NewRootCmd(w *Writers) *cobra.Command {
	flags := &RootFlags{}

	cmd := &cobra.Command{
		Use:     "pig",
		Example: "pig -w 400 -h 300 -f png -t placeholder -c 0xff_ff_ff_ff > output.png",
		Version: "1.0.0",
		Short:   "Generates a placeholder image with the specified width, height, and other options.",
		Long: `placeholder-image-generator(p.i.g.) is a command line application that generates placeholder images.
This application can generate images by specifying width, height, format, background color, and text.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set default text
			if flags.Text == "" {
				flags.Text = fmt.Sprintf("%d x %d", flags.Width, flags.Height)
			}

			return Root(w, flags)
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().Uint16VarP(&flags.Width, "width", "w", 0, "image width pixel")
	cmd.Flags().Uint16VarP(&flags.Height, "height", "h", 0, "image height pixel")
	cmd.Flags().StringVarP(&flags.Format, "format", "f", "jpg", "image format")
	cmd.Flags().StringVarP(&flags.Text, "text", "t", "", "text to be inserted in the image (default \"{width}x{height}\")")
	cmd.Flags().Uint32VarP(&flags.Color, "color", "c", 0x00_00_00_ff, "backgournd color code (RGBA)")

	cmd.Flags().Bool("help", false, "help for "+cmd.Name())
	_ = cmd.Flags().SetAnnotation("help", cobra.FlagSetByCobraAnnotation, []string{"true"})

	cmd.MarkFlagRequired("width")
	cmd.MarkFlagRequired("height")

	return cmd
}

func Root(w *Writers, f *RootFlags) error {
	img, err := imager.Generate(imager.Input{
		Width:  f.Width,
		Height: f.Height,
		Format: f.Format,
		Text:   f.Text,
		Color:  f.Color,
	})
	if err != nil {
		return err
	}

	switch f.Format {
	case imager.Jpg:
		err = jpeg.Encode(w.Out, img, &jpeg.Options{Quality: 70})
	case imager.Png:
		err = png.Encode(w.Out, img)
	case imager.Gif:
		err = gif.Encode(w.Out, img, &gif.Options{})
	}

	return err
}
