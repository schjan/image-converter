package crop

import (
	"errors"
	"image"
)

type SimpleCropper struct {
}

func (c *SimpleCropper) Crop(img image.Image, height, width int) (image.Image, error) {
	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	minX := img.Bounds().Min.X
	maxX := img.Bounds().Max.X
	minY := img.Bounds().Min.Y
	maxY := img.Bounds().Max.Y

	X := maxX - minX
	Y := maxY - minY

	if X < width {
		return nil, errors.New("image has to less width for cropping")
	}
	if Y < height {
		return nil, errors.New("image has to less height for cropping")
	}

	cropMinX := minX + (X / 2) - (width / 2)
	cropMaxX := minX + (X / 2) + (width / 2)
	cropMinY := minY + (Y / 2) - (height / 2)
	cropMaxY := minY + (Y / 2) + (height / 2)

	return img.(SubImager).SubImage(image.Rect(cropMinX, cropMinY, cropMaxX, cropMaxY)), nil
}
