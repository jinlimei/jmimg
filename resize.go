package jmimg

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/draw"
)

func (itu *ImageToUpload) resize(maxWidth, maxHeight int) error {
	mimeType := itu.MimeType()

	bounds := itu.OriginalImage.Bounds()

	// We're going to skip if the w/h of the image is less than our
	// maximums, otherwise we'll be scaling upwards and frankly that
	// is cursed.
	if bounds.Dx() < maxWidth && bounds.Dy() < maxHeight {
		return nil
	}

	var (
		src image.Image
		err error
	)

	switch mimeType {
	case "image/jpeg":
		src, err = jpeg.Decode(itu.Converted)
	case "image/png":
		src, err = png.Decode(itu.Converted)
	case "image/gif":
		src, err = gif.Decode(itu.Converted)
	default:
		return ErrEncodingNotSupported
	}

	if err != nil {
		return err
	}

	if src == nil {
		return fmt.Errorf("src is nil")
	}

	newWidth, newHeight := itu.getResizeHW(src, maxWidth, maxHeight)

	dst := image.NewNRGBA(image.Rect(0, 0, newWidth, newHeight))

	draw.CatmullRom.Scale(dst, dst.Rect, src, bounds, draw.Over, nil)

	buf := new(bytes.Buffer)

	switch mimeType {
	case "image/jpeg":
		err = jpeg.Encode(buf, dst, &jpeg.Options{Quality: 100})
	case "image/png":
		err = png.Encode(buf, dst)
	case "image/gif":
		err = gif.Encode(buf, dst, nil)
	default:
		return ErrEncodingNotSupported
	}

	if err != nil {
		return err
	}

	itu.Converted = buf
	itu.DidResize = true

	return nil
}

func (itu *ImageToUpload) getResizeHW(src image.Image, maxWidth, maxHeight int) (int, int) {
	var (
		maxWH = src.Bounds().Max
		ratio = min(float64(maxWidth)/float64(maxWH.X), float64(maxHeight)/float64(maxWH.Y))
	)

	newW, newH := int(float64(maxWH.X)*ratio), int(float64(maxWH.Y)*ratio)

	fmt.Printf("resized: %d x %d\n", newW, newH)
	return newW, newH
}
