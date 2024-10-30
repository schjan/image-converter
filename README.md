# Image Converter [![Build Status](https://travis-ci.org/schjan/image-converter.svg?branch=master)](https://travis-ci.org/schjan/image-converter)

Example code for my _Go as a Tooling Language for the Cloud_ talk at DevOps Gathering 2019.

To install or update run `go install github.com/schjan/image-converter/cmd/image-converter` after that `image-converter` command is available
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
