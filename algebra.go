package main

import(
	"fmt"
)

func addVector(v1[]float32, v2[]float32){
    rsum :=[]float32{}
	for i := 0; i < len(v1); i++ {
		//fmt.Println(v1[i])
		rsum = append(rsum, (v1[i] + v2[i]))

	}

	fmt.Println(rsum)
}

func subVector(v1[]float32, v2[]float32){
	subResult := []float32{}

	for i:=0; i < len(v1); i++{
		subResult = append(subResult, v1[i] - v2[i])
	}

	fmt.Println(subResult)
}


func main() {

	v1 :=[]float32{1, 2, 3}
	v2 :=[]float32{4, 5, 6}
	addVector(v1, v2)
	subVector(v1, v2)


}