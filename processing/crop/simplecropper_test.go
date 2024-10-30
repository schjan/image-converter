package crop

import (
	"image"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleCropperImageSize(t *testing.T) {
	// Setup
	file, err := os.Open("./../../assets/snow-1920.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	cropper := SimpleCropper{}

	// Test
	result, err := cropper.Crop(img, 400, 300)

	// Assertions
	// For more straightforward (JUnit like) assertions for example use github.com/stretchr/testify
	if err != nil {
		t.Fatalf("Error executing Crop: %+v", err)
	}

	width := result.Bounds().Max.X - result.Bounds().Min.X
	if width != 300 {
		t.Fatalf("Image width should be 300 but was %v", width)
	}

	height := result.Bounds().Max.Y - result.Bounds().Min.Y
	if height != 400 {
		t.Fatalf("Image width should be 400 but was %v", height)
	}
}

func TestSimpleCropperImageSizeTestify(t *testing.T) {
	// Setup
	file, err := os.Open("./../../assets/snow-1920.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	cropper := SimpleCropper{}

	// Test
	result, err := cropper.Crop(img, 400, 300)

	// Assertions
	require.NoError(t, err)

	width := result.Bounds().Max.X - result.Bounds().Min.X
	height := result.Bounds().Max.Y - result.Bounds().Min.Y
	assert.Equal(t, 300, width)
	assert.Equal(t, 400, height)
}
