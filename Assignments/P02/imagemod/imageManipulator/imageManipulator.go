// imagemod/imageManipulator/imageManipulator.go

package imageManipulator

import (
	"github.com/fogleman/gg"
)

// ImageManipulator represents an image manipulation tool.
type ImageManipulator struct {
	Image *gg.Context
}

// NewImageManipulator creates a new ImageManipulator instance.
func NewImageManipulator(width, height int) *ImageManipulator {
	img := gg.NewContext(width, height)
	return &ImageManipulator{Image: img}
}

// SaveToFile saves the manipulated image to a file.
func (im *ImageManipulator) SaveToFile(filename string) error {
	return im.Image.SavePNG(filename)
}

// DrawRectangle draws a rectangle on the image.
func (im *ImageManipulator) DrawRectangle(x, y, width, height float64) {
	im.Image.DrawRectangle(x, y, width, height)
	im.Image.Stroke()
}
