# pig

placeholder-image-generator(p.i.g.) is a command line application that generates placeholder images.  
This application can generate images by specifying width, height, format, background color, and text.

```
pig [flags]
```

## Examples

```
pig -w 400 -h 300 -f png -t placeholder -c 0xff_ff_ff_ff > output.png
```

## Options

```
  -w, --width uint16    image width pixel
  -h, --height uint16   image height pixel
  -f, --format string   image format (default "jpg")
  -t, --text string     text to be inserted in the image (default "{width}x{height}")
  -c, --color uint32    backgournd color code (RGBA) (default 255)
      --help            help for pig
```
