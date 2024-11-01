package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"slices"
	"strings"

	"github.com/schjan/image-converter/processing"
	"github.com/schjan/image-converter/processing/crop"
)

// see explanation of flags with -help flag
var sourceDirectory = flag.String("source", "./", "directory containing source image files")
var outputDirectory = flag.String("out", "./out", "directory to save converted images")
var smartCropper = flag.Bool("smart", false, "if set use smart cropper")

func main() {
	flag.Parse()

	if err := mainErr(); err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func mainErr() error {
	sourceFiles, err := imagesInDir(*sourceDirectory)
	if err != nil {
		return fmt.Errorf("could not read source directory %s: %w", *sourceDirectory, err)
	}

	// choose which implementation of Cropper to use with -smart flag
	var cropper crop.Cropper
	if *smartCropper {
		log.Print("This software uses advanced cropping technologies 💥")
		cropper = crop.NewSmart()
	} else {
		cropper = &crop.SimpleCropper{}
	}

	processor, err := processing.New(cropper, runtime.NumCPU())
	if err != nil {
		return fmt.Errorf("could not create image processor: %w", err)
	}

	for _, srcFilename := range sourceFiles {
		outFilename := path.Join(*outputDirectory, strings.TrimSuffix(filepath.Base(srcFilename), filepath.Ext(srcFilename))+".png")

		processor.AddToQueue(srcFilename, outFilename)
	}

	processor.StopAndWaitFinished()

	return nil
}

var supportedExtensions = []string{".png", ".jpg", ".jpeg"}

func imagesInDir(directory string) ([]string, error) {
	dirFiles, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("could not read source directory: %w", err)
	}

	var files []string
	for _, fileInfo := range dirFiles {
		if fileInfo.IsDir() {
			continue
		}

		ext := filepath.Ext(fileInfo.Name())
		if slices.Contains(supportedExtensions, ext) {
			files = append(files, path.Join(directory, fileInfo.Name()))
		}
		// Before slices package:
		//for _, supportedExt := range supportedExtensions {
		//	 if ext == supportedExt {
		//		 files = append(files, path.Join(directory, fileInfo.Name()))
		//	 }
		//}
	}

	return files, nil
}
