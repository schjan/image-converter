package crop

import (
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"github.com/muesli/smartcrop/options"
	"image"
)

type smartCropper struct {
	analyzer smartcrop.Analyzer
	resizer  options.Resizer
}

func NewSmart() *smartCropper {
	rz := nfnt.NewDefaultResizer()

	return &smartCropper{resizer: rz, analyzer: smartcrop.NewAnalyzer(rz)}
}

func (c *smartCropper) Crop(img image.Image, height, width int) (image.Image, error) {
	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	topCrop, err := c.analyzer.FindBestCrop(img, width, height)
	if err != nil {
		return nil, err
	}

	img = img.(SubImager).SubImage(topCrop)

	return c.resizer.Resize(img, uint(width), uint(height)), nil
}
