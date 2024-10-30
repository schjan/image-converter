package processing

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/schjan/image-converter/processing/crop"
)

type processor struct {
	cropper     crop.Cropper
	jobChan     chan jobOptions
	stoppedChan chan struct{}
	workerCount int
}

type jobOptions struct {
	sourcePath string
	outputPath string
}

func New(cropper crop.Cropper, workers int) (*processor, error) {
	if workers == 0 || workers > 20 {
		return nil, errors.New("invalid number of workers")
	}

	p := &processor{
		jobChan:     make(chan jobOptions, 40),
		stoppedChan: make(chan struct{}),
		cropper:     cropper,
		workerCount: workers,
	}

	for i := 0; i < workers; i++ {
		go p.worker()
	}

	return p, nil
}

func (p *processor) AddToQueue(sourcePath string, destinationPath string) {
	p.jobChan <- jobOptions{
		sourcePath: sourcePath,
		outputPath: destinationPath,
	}
}

func (p *processor) worker() {
	for {
		job, ok := <-p.jobChan
		if !ok {
			p.stoppedChan <- struct{}{}
			return
		}

		err := p.processJob(job)
		if err != nil {
			log.Printf("Error processing image %v", job.sourcePath)
		} else {
			log.Printf("Successfull processed image %v", job.sourcePath)
		}
	}
}

func (p *processor) processJob(job jobOptions) error {
	file, err := os.Open(job.sourcePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// first try to create output directory, there won't be an error if directory already exists
	err = os.MkdirAll(filepath.Dir(job.outputPath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not create path structure for destination file '%s': %w", job.outputPath, err)
	}

	outFile, err := os.Create(job.outputPath)
	if err != nil {
		return fmt.Errorf("could not create file '%s': %w", job.outputPath, err)
	}
	defer outFile.Close()

	err = p.processImage(outFile, file)
	if err != nil {
		return fmt.Errorf("converting '%s' failed: %w", job.sourcePath, err)
	}

	// with defer outFile.Close() we can't see any errors created while flushing buffers etc.
	err = outFile.Close()
	if err != nil {
		return fmt.Errorf("error saving file '%s': %w", job.outputPath, err)
	}

	return nil
}

func (p *processor) processImage(writer io.Writer, reader io.Reader) error {
	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	img, err = p.cropper.Crop(img, 300, 300)
	if err != nil {
		return err
	}

	err = png.Encode(writer, img)
	if err != nil {
		return err
	}

	return nil
}

func (p *processor) StopAndWaitFinished() {
	// closes the channel after the last item is received
	close(p.jobChan)

	// every worker should send to stoppedChan before shutting down
	workersStopped := 0
	for workersStopped != p.workerCount {
		<-p.stoppedChan
		workersStopped++
	}
}
