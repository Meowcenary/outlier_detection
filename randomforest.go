package main

import (
	"fmt"
	"math/rand"

	randomforest "github.com/malaschitz/randomForest"
	"github.com/petar/GoMNIST"
)

func main() {
	// initialize randomness for isolation forest
	rand.Seed(1)
	// number of trees for forest
	TREES := 1000
	// number of pictures in the MNIST data set
	size := 60000
	// The number of pixels for a single picture from the MNIST data set
	xsize := 28 * 28

	// labels is an array of []GoMNIST.Label which are uint8
	// the value of GoMNIST.Label is the number that the image contains
	// e.g 2 would be a picture of a 2
	labels, err := GoMNIST.ReadLabelFile("data/train-labels-idx1-ubyte.gz")
	if err != nil {
		panic(err)
	}

	// imgs has type []GoMNIST.RawImage
	// GoMNIST.RawImage is type []byte which holds the pixel intensities of an image.
	// 255 is foreground (black), 0 is background (white).
	_, _, imgs, err := GoMNIST.ReadImageFile("data/train-images-idx3-ubyte.gz")
	if err != nil {
		panic(err)
	}

	// Size is the number of images in the set
	// The labels and images must match this
	if len(labels) != size || len(imgs) != size {
		panic("Wrong size")
	}
	// Train

	// Instatiate a randomForest  Forest object
	forest := randomforest.Forest{}
	// Instantiate a 2D slice of float64 that is the size of the number of images
	x := make([][]float64, size)
	// Int slice for labels
	l := make([]int, size)

	for i := 0; i < size; i++ {
		x[i] = make([]float64, xsize)

		for j := 0; j < xsize; j++ {
			// Remember, imgs[i] is type []GoMNIST.RawImage and RawImage is is type []byte
			// where the bytes are from 0-255. Here the byte is cast to a float64 type so it
			// can be analyzed by the randomForest package
			x[i][j] = float64(imgs[i][j])
			// labels[i] is a value between 0-9, but is stored as a ubyte8 so it must be cast to int here
			l[i] = int(labels[i])
		}
	}

	// Now that the pixel data has been converted to float64 and the labels have been convertd to ints
	// tehy can be stored into a ForestData object
	forest.Data = randomforest.ForestData{X: x, Class: l}
	forest.MaxDepth = 30
	forest.Train(TREES)

	// //ISOLATION
	isolations, mean, stddev := forest.IsolationForest()
	for i, d := range isolations {
		if d < (mean - 1.6*stddev) {
			fmt.Println(i, (d-mean)/stddev)
		}
	}
	//
}
