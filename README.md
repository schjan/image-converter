# Image Converter [![Build Status](https://travis-ci.org/schjan/image-converter.svg?branch=master)](https://travis-ci.org/schjan/image-converter)

Example code for my _Java on the Go - Go als Hilfsmittel für den Java-Veteranen_ talk at JavaLand 2018.

To install or update run `go get -u github.com/schjan/image-converter/...` after that `image-converter` command is available
from bash / cmd.


```
> image-converter -help

Usage of image-converter:
  -out string
        directory to save converted images (default "./out")
  -smart
        if set use smart cropper (github.com/muesli/smartcrop)
  -source string
        directory containing source image files (default "./")
```