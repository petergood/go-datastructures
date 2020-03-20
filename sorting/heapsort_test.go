package sorting

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{2, 5, 1, 7, 4, 2, 10}
	expectedAsc := []int{1, 2, 2, 4, 5, 7, 10}
	expectedDesc := []int{10, 7, 5, 4, 2, 2, 1}

	resAsc := Sort(arr, Asc)
	resDesc := Sort(arr, Desc)

	if !reflect.DeepEqual(resAsc, expectedAsc) {
		t.Errorf("Invalid asc array %v", resAsc)
	}

	if !reflect.DeepEqual(resDesc, expectedDesc) {
		t.Errorf("invalid desc array %v", resDesc)
	}
}
