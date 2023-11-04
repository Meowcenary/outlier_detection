package main

import (
	// "fmt"
	// "math/rand"
  //
	// randomForest "github.com/malaschitz/randomForest"
	"strconv"
	"github.com/petar/GoMNIST"
)

func main() {
	// Read both the MNIST training and testing data using GoMNIST
	// train and test are assigned to pointers to Set types
	// where set is defined within the package GoMNIST as:
	// Set {
	//   NRow int
	//   NCol nit
	//   Images []RawImage , RawImage is a []byte
	//   Labels []Label , Label is a uint8 , digit label 0 to 9
	// }
	train, _, err := GoMNIST.Load("./data")

	if err != nil {
		panic(err)
	}

	ReshapeImageData(train)
	// Outlier detection with isolated forest
	// randomForest needs the data in the form [][]float64
	// xData := [][]float64{}
	// yData := []int{}
	// forest := randomForest.Forest{}
	// forest.Data = randomforest.ForestData{X: xData, Class: yData}
	// replace 1000 with whatever number of values there are to compare
	// forest.Train(1000)
}

// The images are stored as []byte, but randomForest.Forest needs [][]float64
func ReshapeImageData(set *GoMNIST.Set) [][]float64 {
	setObject := *set
	images := setObject.Images
	imageCount := len(setObject.Images)
	reshapedData := [][]float64{}

	for i := 0; i < imageCount; i++ {
		imageDataAsFloat := strconv.ParseFloat(string(images[i]))

		append(reshapedData, imageDataAsFloat)
	}

	return reshapedData
}
