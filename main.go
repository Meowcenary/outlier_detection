package main

import (
	// "fmt"
	// "math/rand"
  //
	// randomForest "github.com/malaschitz/randomForest"
	"github.com/Meowcenary/outlier_detection/compressed_data_reader"
)

func main() {
	compressed_data_reader.Decompress("DERP")
	// xData := [][]float64{}
	// yData := []int{}
  //
	// for i := 0; i < 1000; i++ {
	// 	x := []float64{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}
	// 	y := int(x[0] + x[1] + x[2] + x[3])
	// 	xData = append(xData, x)
	// 	yData = append(yData, y)
	// }
  //
	// forest := randomForest.Forest{}
	// forest.Data = randomForest.ForestData{X: xData, Class: yData}
	// forest.Train(1000)
  //
	// //test
	// fmt.Println("Vote", forest.Vote([]float64{0.1, 0.1, 0.1, 0.1}))
	// fmt.Println("Vote", forest.Vote([]float64{0.9, 0.9, 0.9, 0.9}))
}
