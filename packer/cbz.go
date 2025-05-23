package packer

import (
	"archive/zip"
	"errors"
	"fmt"
	"os"

	"github.com/elboletaire/manga-downloader/downloader"
)

// ArchiveCBZ archives the given files into a CBZ file
func ArchiveCBZ(filename string, files []*downloader.File, progress func(page, progress int)) error {
	if len(files) == 0 {
		return errors.New("no files to pack")
	}
	buff, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	defer buff.Close()
	w := zip.NewWriter(buff)

	for i, file := range files {
		f, err := w.Create(fmt.Sprintf("%03d.jpg", i))
		if err != nil {
			return err
		}
		if _, err = f.Write(file.Data); err != nil {
			return err
		}
		progress(1, 0) // Report progress by single page increments
	}

	err = w.Close()
	return err
}
