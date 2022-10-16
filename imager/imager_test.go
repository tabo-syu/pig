package imager

import (
	_ "embed"
	"errors"
	"image"
	"image/color"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		input Input
	}
	type testcase struct {
		name    string
		args    args
		want    *image.RGBA
		wantErr error
	}

	tests := []testcase{
		{
			name:    "Test image width, height validation",
			args:    args{Input{Width: 0, Height: 0, Format: "jpg"}},
			want:    nil,
			wantErr: ErrInvalidImageSize,
		},
		{
			name:    "Test image format validation",
			args:    args{Input{Width: 100, Height: 100, Format: "hoge"}},
			want:    nil,
			wantErr: ErrInvalidImageFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.input)
			if err != nil && !errors.Is(tt.wantErr, err) {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fillColor(t *testing.T) {
	type args struct {
		img *image.RGBA
		c   color.RGBA
	}
	type testcase struct {
		name string
		args args
		want color.RGBA
	}
	tests := []testcase{
		{
			name: "Test fill red",
			args: args{
				img: image.NewRGBA(image.Rect(0, 0, 400, 300)),
				c:   color.RGBA{255, 0, 0, 255},
			},
			want: color.RGBA{255, 0, 0, 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fillColor(tt.args.img, tt.args.c)
			size := tt.args.img.Bounds().Size()
			for h := 0; h < size.Y; h++ {
				for w := 0; w < size.X; w++ {
					got := tt.args.img.At(w, h)
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("fillColor() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}

func Test_colorCodeToRGBA(t *testing.T) {
	type args struct {
		code uint32
	}
	type testcase struct {
		name  string
		args  args
		wantR uint8
		wantG uint8
		wantB uint8
		wantA uint8
	}
	tests := []testcase{
		{
			name: "Test red",
			args: args{
				code: 0xff_00_00_ff,
			},
			wantR: 255,
			wantG: 0,
			wantB: 0,
			wantA: 255,
		},
		{
			name: "Test green",
			args: args{
				code: 0x00_ff_00_ff,
			},
			wantR: 0,
			wantG: 255,
			wantB: 0,
			wantA: 255,
		},
		{
			name: "Test blue",
			args: args{
				code: 0x00_00_ff_ff,
			},
			wantR: 0,
			wantG: 0,
			wantB: 255,
			wantA: 255,
		},
		{
			name: "Test black",
			args: args{
				code: 0x00_00_00_ff,
			},
			wantR: 0,
			wantG: 0,
			wantB: 0,
			wantA: 255,
		},
		{
			name: "Test white",
			args: args{
				code: 0xff_ff_ff_ff,
			},
			wantR: 255,
			wantG: 255,
			wantB: 255,
			wantA: 255,
		},
		{
			name: "Test color",
			args: args{
				code: 0x24_42_17_64,
			},
			wantR: 36,
			wantG: 66,
			wantB: 23,
			wantA: 100,
		},
		{
			name: "Test missing color code",
			args: args{
				code: 0xff_ff_ff,
			},
			wantR: 0,
			wantG: 255,
			wantB: 255,
			wantA: 255,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB, gotA := colorCodeToRGBA(tt.args.code)
			if gotR != tt.wantR {
				t.Errorf("colorCodeToRGBA() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotG != tt.wantG {
				t.Errorf("colorCodeToRGBA() gotG = %v, want %v", gotG, tt.wantG)
			}
			if gotB != tt.wantB {
				t.Errorf("colorCodeToRGBA() gotB = %v, want %v", gotB, tt.wantB)
			}
			if gotA != tt.wantA {
				t.Errorf("colorCodeToRGBA() gotA = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}
