package crop

import (
	"bytes"
	"image"
	"os"
	"testing"
)

var result image.Image

func BenchmarkCroppers(b *testing.B) {
	bts, err := os.ReadFile("./../../assets/snow-1920.jpg")
	if err != nil {
		b.Fatal(err)
	}

	b.Run("Simple", func(b *testing.B) {
		benchmarkCropper(b, &SimpleCropper{}, bts)
	})
	b.Run("Smart", func(b *testing.B) {
		benchmarkCropper(b, NewSmart(), bts)
	})
}

func benchmarkCropper(b *testing.B, cropper Cropper, bts []byte) {
	var res image.Image

	for i := 0; i < b.N; i++ {
		img, _, _ := image.Decode(bytes.NewReader(bts))
		res, _ = cropper.Crop(img, 200, 200)
	}

	result = res
}
