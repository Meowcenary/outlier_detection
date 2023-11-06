package main

import (
	"fmt"
	"math/rand"
	"sort"

	randomforest "github.com/malaschitz/randomForest"
	"github.com/petar/GoMNIST"
)

func main() {
	rand.Seed(1)

	// loop 10 times
	for l := 0; l < 10; l++ {
		TREES := 100000
		size := 60000
		// 28x28 pixel images
		xsize := 28 * 28

		// Read label file
		labels, err := GoMNIST.ReadLabelFile("data/train-labels-idx1-ubyte.gz")
		if err != nil {
			panic(err)
		}

		// Read image file
		_, _, imgs, err := GoMNIST.ReadImageFile("data/train-images-idx3-ubyte.gz")
		if err != nil {
			panic(err)
		}

		// Train
		// 2D array of float64 all initialized to 0
		x := make([][]float64, 0)
		// Blank map of int to int
		mp := map[int]int{}

		for i := 0; i < size; i++ {

			// if labels[i] to int == current loop var l
			if int(labels[i]) == l {
				// make an array of float64
				xx := make([]float64, xsize)

				// and fill it with image data
				for j := 0; j < xsize; j++ {
					xx[j] = float64(imgs[i][j])
				}

				x = append(x, xx)
				mp[len(x)] = i
			}
		}

		fmt.Println(len(x), "x", len(x[0]))
		//
		forest := randomforest.IsolationForest{X: x}
		forest.Train(TREES)
		// results
		a := make([]int, len(x))
		for i := 0; i < len(x); i++ {
			a[i] = i
		}
		sort.Slice(a, func(i, j int) bool {
			ai := forest.Results[a[i]]
			aj := forest.Results[a[j]]
			mi := float64(ai[1]) / (float64(ai[0]) + 0.00001)
			mj := float64(aj[1]) / (float64(aj[0]) + 0.00001)
			return mi < mj
		})
		fmt.Println("-------------", l)
		for i := 0; i < 10; i++ {
			fmt.Printf("%3d %5d %3d %3d %5.2f\n", i, mp[a[i]], forest.Results[a[i]][0], forest.Results[a[i]][1],
				float64(forest.Results[a[i]][1])/(float64(forest.Results[a[i]][0])+0.00001))
		}
	}
}
