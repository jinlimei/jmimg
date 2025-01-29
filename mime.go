package jmimg

import (
	"io"
	"net/http"
)

func (itu *ImageToUpload) determineMimeType() error {
	first512 := make([]byte, 512)

	_, err := itu.Original.Read(first512)
	if err != nil {
		return err
	}

	_, err = itu.Original.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	itu.OriginalMimeType = http.DetectContentType(first512)
	return nil
}
