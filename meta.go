package jmimg

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

func (itu *ImageToUpload) parseImageMetadata() error {
	stat, err := itu.Original.Stat()
	if err != nil {
		return err
	}

	itu.Stat = stat

	var src image.Image

	switch itu.OriginalMimeType {
	case "image/jpeg":
		src, err = jpeg.Decode(itu.Original)
	case "image/png":
		src, err = png.Decode(itu.Original)
	case "image/gif":
		src, err = gif.Decode(itu.Original)
	}

	if err != nil {
		return err
	}

	itu.OriginalImage = src

	// Reset original to seek to beginning
	// Necessary because .Decode will read the stream to its end.
	_, err = itu.Original.Seek(0, 0)

	if err != nil {
		return err
	}

	itu.Converted = itu.Original

	return nil
}
