package compressed_data_reader

import (
	"compress/gzip"
	"io"
	"path"
	"os"
)

func Decompress(filepath string) {
	// Open the gzipped file
	gzippedFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer gzippedFile.Close()

	// Create a new gzip reader
	gzipReader, err := gzip.NewReader(gzippedFile)
	defer gzipReader.Close()

	ext := path.Ext(filepath)
	uncompressedFilePath := filepath[0:len(filepath)-len(ext)]

	// Create a new file to hold the uncompressed data
	uncompressedFile, err := os.Create(uncompressedFilePath)
	if err != nil {
			panic(err)
	}
	defer uncompressedFile.Close()

	// Copy the contents of the gzip reader to the new file
	_, err = io.Copy(uncompressedFile, gzipReader)
	if err != nil {
			panic(err)
	}
}
