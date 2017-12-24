package main

import (
	"fmt"
)
const N int = 6
var a = [N]int{0,1,2,3,4,5}

func main(){
	t := 3
	var i,j,k int
	for i, j = 0, N - 1;i <= j;{
		if a[(i + j) / 2] > t{
			j = i + (j - i) / 2
		}else if a[(i +j) / 2] < t{
			i = j - (j - i) / 2
		}else{
			k = (i + j) / 2
			break
		}
	}
	
	if a[k] != t{
		k = -1
	}
	fmt.Println(k)	
}
