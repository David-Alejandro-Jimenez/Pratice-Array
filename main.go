package main

import (
	"fmt"
)

//La funcion ArraysPrueba se encarga de alternar valores de dos arrays y poner el resultado en un array vacio
func ArraysPrueba(array1, array2 []int) (resultado []int) {
	for i := range array1 {
		resultado = append(resultado, array1[i], array2[i])
	}

	fmt.Println(resultado[0:9])

	return resultado 
	
}

func main() {
	ArraysPrueba([]int{1,2,3,4,5}, []int{0,9,8,7,0})
}