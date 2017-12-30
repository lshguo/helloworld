package main

import (
	"fmt"
)

var aa = [][]int{
	{6,5,4,3,2,1},
	{1,2,3,4,5,6},
	{1,1,1,1,1,1},
	{1,1},
	{1,2},
	{2,1},
	{1},
	{}}

func quickSort(a []int) {
	if len(a) <= 1{
		return
	}

	head, tail := 0, len(a) - 1
	mid := a[head]

	for head < tail{
		if a[head + 1] <= mid{
			a[head], a[head + 1] = a[head + 1], a[head]
			head++
		}else{
			a[head + 1], a[tail] = a[tail], a[head + 1]
			tail--
		}
	}
	//go语法里，slice终点不可以越界
	quickSort(a[:head])
	//go语法里，slice起点可以越界
	quickSort(a[head+1:])
}
func main(){
	for _, a := range aa{
		quickSort(a)
		fmt.Println(a)
	}
}





