package main

import (
	"fmt"
)

//Creamos en una funcion los arrays y su tipo de dato
func ArraysPrueba(Array1 []int, Array2 []int) (Resultado []int) {
	//Creamos un ciclo for para que alterne los 2 arrays
	for i := range Array1 {
		Resultado = append(Resultado, Array1[i], Array2[i],)
	}
	//Imprimimos El Resultado.
	fmt.Println(Resultado[0:9])
	//retornamos el resultado del ciclo for
	return Resultado
	
}

func main() {
	//Por ultimo llamamos a la fucnion ArraysPrueba y le asignamos los valores al Array
	ArraysPrueba([]int{1,2,3,4,5}, []int{0,9,8,7,0})
}