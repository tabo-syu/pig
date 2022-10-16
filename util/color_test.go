package util

import (
	"testing"
)

func TestColorCodeToRGBA(t *testing.T) {
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
			gotR, gotG, gotB, gotA := ColorCodeToRGBA(tt.args.code)
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

func TestCalcColorFromBGColor(t *testing.T) {
	type args struct {
		ir uint8
		ig uint8
		ib uint8
	}
	type testcase struct {
		name  string
		args  args
		wantR uint8
		wantG uint8
		wantB uint8
	}
	tests := []testcase{
		{
			name: "Textcolor will be black",
			args: args{
				ir: 0xff,
				ig: 0xff,
				ib: 0xff,
			},
			wantR: 0,
			wantG: 0,
			wantB: 0,
		},
		{
			name: "Textcolor will be white",
			args: args{
				ir: 0,
				ig: 0,
				ib: 0,
			},
			wantR: 0xff,
			wantG: 0xff,
			wantB: 0xff,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB := CalcColorFromBGColor(tt.args.ir, tt.args.ig, tt.args.ib)
			if gotR != tt.wantR {
				t.Errorf("CalcColorFromBGColor() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotG != tt.wantG {
				t.Errorf("CalcColorFromBGColor() gotG = %v, want %v", gotG, tt.wantG)
			}
			if gotB != tt.wantB {
				t.Errorf("CalcColorFromBGColor() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
