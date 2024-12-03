package main

import(
	"fmt"
)

func addVector(v1[]float32, v2[]float32)[]float32{
    sumResult :=[]float32{}
	for i := 0; i < len(v1); i++ {
		//fmt.Println(v1[i])
		sumResult = append(sumResult, (v1[i] + v2[i]))

	}

	//fmt.Println(rsum)
	return sumResult
}

func subVector(v1[]float32, v2[]float32)[]float32{
	subResult := []float32{}

	for i:=0; i < len(v1); i++{
		subResult = append(subResult, v1[i] - v2[i])
	}

	//fmt.Println(subResult)
	return subResult
}


func main() {

	v1 :=[]float32{1, 2, 3}
	v2 :=[]float32{4, 5, 6}

	fmt.Println(addVector(v1, v2))
	fmt.Println(subVector(v1, v2))


}