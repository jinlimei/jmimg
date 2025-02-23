package jmimg

import (
	"image"
	"io"
	"os"
)

type ImageToUpload struct {
	FileName string

	Original         *os.File
	OriginalMimeType string

	Stat          os.FileInfo
	OriginalImage image.Image

	Converted io.Reader

	DidConvert       bool
	DidResize        bool
	DidCleanMetadata bool
}

func (itu *ImageToUpload) MimeType() string {
	if itu.DidConvert {
		return "image/jpeg"
	}

	return itu.OriginalMimeType
}

func (itu *ImageToUpload) Reader() io.Reader {
	if itu.DidConvert || itu.DidResize || itu.DidCleanMetadata {
		return itu.Converted
	}

	return itu.Original
}

func NewImageUpload(fileName string, file *os.File) (*ImageToUpload, error) {
	itu := &ImageToUpload{
		FileName:  fileName,
		Original:  file,
		Converted: file,
	}

	err := itu.determineMimeType()

	if err != nil {
		return nil, err
	}

	err = itu.parseImageMetadata()

	if err != nil {
		return nil, err
	}

	return itu, nil
}
