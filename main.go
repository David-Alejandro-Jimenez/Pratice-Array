package main

import (
	"fmt"
)

//Creamos en una funcion los arrays y su tipo de dato
func ArraysPrueba(array1, array2 []int) (resultado []int) {
	//Creamos un ciclo for para que alterne los 2 arrays
	for i := range array1 {
		resultado = append(resultado, array1[i], array2[i])
	}
	//Imprimimos El Resultado.
	fmt.Println(resultado[0:9])
	//retornamos el resultado del ciclo for
	return resultado 
	
}

func main() {
	//Por ultimo llamamos a la fucnion ArraysPrueba y le asignamos los valores al Array
	ArraysPrueba([]int{1,2,3,4,5}, []int{0,9,8,7,0})
}