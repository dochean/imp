package imp

import (
	"fmt"
	"testing"
)

func TestSortBubble(t *testing.T) {
	data := [][]int{
		{2, 25, 7, 29, 13},
		{5, 4, 3, 2, 1},
	}
	for _, d := range data {
		SortBubble(d)
		fmt.Println(d)
	}
}

func TestSortInsert(t *testing.T) {
	data := [][]int{
		{2, 25, 7, 29, 13},
		{5, 4, 3, 2, 1},
	}
	for _, d := range data {
		SortInsert(d)
		fmt.Println(d)
	}
}

func TestSortSelect(t *testing.T) {
	data := [][]int{
		{2, 25, 7, 29, 13},
		{5, 4, 3, 2, 1},
	}
	for _, d := range data {
		SortSelect(d)
		fmt.Println(d)
	}
}