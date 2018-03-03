package crop

import "image"

type Cropper interface {
	Crop(img image.Image, height, width int) (image.Image, error)
}
