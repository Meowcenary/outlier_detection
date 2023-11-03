package compressed_data_reader

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"path"
	"os"
)

// this function heavily references this script:
//   https://github.com/ThomasWMiller/jump-start-mnist-iforest/blob/main/python/getMNIST.py
func ReadImagesFrom(filepath string) {
	bytes := ReadBytesFrom(filepath)
	// images are 28x28 pixels and pixels in RGB require 3 bytes each to store
	// pixels in black and white only require 1 byte each
	imageSize := 28
	numImages := 60000
	// in the python script referenced above there is a call to read() with 16 passed
	// which seems to be done to skip over the first 16 units
	offset := 16

	for i := 0; i < numImages; i++ {
		startByte := i*imageSize*imageSize
		endByte := (i+1)*imageSize*imageSize

		bytes[startByte:endByte]
	}
}

func ReadBytesFrom(filepath string) []byte {
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		panic(err)
	}

	return bs
}

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
