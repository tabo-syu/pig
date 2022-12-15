# pig

placeholder-image-generator(p.i.g.) is a command line application that generates placeholder images.  
This application can generate images by specifying width, height, format, background color, and text.

```
$ pig [flags]
```

## Examples

```bash
$ pig -w 400 -h 300 -f png -t placeholder -c 0xff_ff_ff_ff > output.png
```

Result:  
![Result image](https://user-images.githubusercontent.com/45633620/207982130-8bb9d310-e57a-44ad-814c-3f027f8bd047.png)

## Options

```
  -w, --width uint16    image width pixel
  -h, --height uint16   image height pixel
  -f, --format string   image format (default "jpg")
  -t, --text string     text to be inserted in the image (default "{width}x{height}")
  -c, --color uint32    backgournd color code (RGBA) (default 255)
      --help            help for pig
```

## Installation

```bash
$ go install github.com/tabo-syu/pig@latest
```
