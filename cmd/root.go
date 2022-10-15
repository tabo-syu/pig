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
	"io"

	"github.com/spf13/cobra"
)

type Writers struct {
	Out io.Writer
	Err io.Writer
}

type RootFlags struct {
	Width    uint
	Height   uint
	Format   string
	Text     string
	Filename string
}

func NewRootCmd(w *Writers) *cobra.Command {
	flags := &RootFlags{}

	cmd := &cobra.Command{
		Use:     "dummy-image-generator",
		Version: "1.0.0",
		Short:   "Generates a dummy image with the specified width, height, and other specified options.",
		Long: `dummy-image-generator is a command line application that generates dummy images.
This application can generate images by specifying width, height, format, file name, and text.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return root(w, flags)
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().UintVarP(&flags.Width, "width", "w", 0, "image width pixel")
	cmd.Flags().UintVarP(&flags.Height, "height", "h", 0, "image height pixel")
	cmd.Flags().StringVarP(&flags.Format, "format", "f", "jpg", "image format")
	cmd.Flags().StringVarP(&flags.Text, "text", "t", "", "text to be inserted in the image (default \"{width}x{height}\")")
	cmd.Flags().StringVarP(&flags.Filename, "filename", "n", "", "filename (default \"{width}x{height}.{format}\")")
	cmd.Flags().Bool("help", false, "help for "+cmd.Name())
	_ = cmd.Flags().SetAnnotation("help", cobra.FlagSetByCobraAnnotation, []string{"true"})

	cmd.MarkFlagsRequiredTogether("width", "height")

	return cmd
}

func root(w *Writers, f *RootFlags) error {
	fmt.Fprintln(
		w.Out,
		f.Width, "x", f.Height, ".", f.Format,
	)

	return nil
}
