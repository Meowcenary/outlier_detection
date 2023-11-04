package main

import (
	// "fmt"
	// "math/rand"
  //
	randomForest "github.com/malaschitz/randomForest"
	"github.com/petar/GoMNIST"
)

func main() {
	// Read the MNIST data using GoMNIST
	train, test, err := Load("./data")

	if err != nil {
		panic(err)
	}

	// TODO: what data is there to even work with here?
	// Outlier detection with isolated forest
	xData := [][]float64{}
	yData := []int{}
	forest := randomForest.Forest{}
	forest.Data = randomforest.ForestData{X: xData, Class: yData}
	// replace 1000 with whatever number of values there are to compare
	forest.Train(1000)
}
