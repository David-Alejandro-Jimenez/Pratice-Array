package main

import (
	"slices"
	"testing"
)

func TestArrays(t *testing.T) {
	var test = []struct {
		numbers1 []int
		numbers2 []int
		result []int
	}{
		{[]int{1,2,3,4,5}, []int{0,9,8,7,0}, []int{1,0,2,9,3,8,4,7,5,0}},
	}

	for _, i := range test {
		var res = slices.Equal(ArraysPrueba(i.numbers1, i.numbers2), i.result)
		if res != true {
			t.Errorf("Practica incorrecta esperabamos una union entre %v y %v pero obtuvimos otra union inesperada %v", i.numbers1, i.numbers2, i.result)
		}
	}
}