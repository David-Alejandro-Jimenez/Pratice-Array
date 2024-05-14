package main

import "testing"

func TestArrays(t *testing.T) {
	numbers1 := []int{1,2,3,4,5}
	numbers2 := []int{0,9,8,7,0}

	result := ArraysPrueba(numbers1, numbers2)
	for i := range numbers1 {
		result = append(result, numbers1[i], numbers2[i])
		if result == nil {
			t.Errorf("Practica incorrecta")
		}
	} 
} 