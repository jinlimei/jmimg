package jmimg

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

var (
	ErrEncodingNotSupported = errors.New("encoding not supported")
)

func (itu *ImageToUpload) changeToJPEG() error {
	var (
		img image.Image
		err error
	)

	switch itu.OriginalMimeType {
	// We're going to skip re-encoding jpegs!
	case "image/jpeg":
		itu.Converted = itu.Original
		return nil

	case "image/png":
		img, err = png.Decode(itu.Original)
	case "image/gif":
		img, err = gif.Decode(itu.Original)
	default:
		return ErrEncodingNotSupported
	}

	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 100})

	if err != nil {
		return err
	}

	itu.Converted = buf
	itu.DidConvert = true

	return nil
}
