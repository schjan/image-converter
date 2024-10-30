package crop

import (
	"bytes"
	"image"
	"io"
	"os"
	"testing"
)

var result image.Image

func BenchmarkCroppers(b *testing.B) {
	file, err := os.Open("./../../assets/snow-1920.jpg")
	if err != nil {
		b.Fatal(err)
	}
	defer file.Close()

	bts, err := io.ReadAll(file)
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
