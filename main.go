/*
Copyright © 2022 tabosuke

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
package main

import (
	"os"

	"github.com/tabo-syu/dummy-image-generator/cmd"
)

type exitCode int

const (
	exitOK    exitCode = 0
	exitError exitCode = 1
)

func main() {
	code := exec()
	os.Exit(int(code))
}

func exec() exitCode {
	writers := &cmd.Writers{Out: os.Stdout, Err: os.Stderr}
	rootCmd := cmd.NewRootCmd(writers)
	if err := rootCmd.Execute(); err != nil {
		return exitError
	}

	return exitOK
}
