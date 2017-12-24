package main

import (
	"fmt"
)
var cnt int = 0
const M int = 4
const N int = 4
var a = [M][N]int {
	{0,0,1,0},
	{0,0,1,0},
	{0,0,1,0},
	{0,0,1,0},
}
func nextStep(i, j int){
	if M - 1 == i  && N - 1 == j{
		cnt++
		return 
	}

	//right
	if j != N - 1  && a[i][j + 1] != 1{
		nextStep(i, j + 1)
	}
	//down
        if i != M - 1  && a[i + 1][j] != 1{
                nextStep(i + 1, j)
        }
}
func main(){
	//fmt.Println(M,N)
	//fmt.Println(a[2][2])	
	nextStep(0, 0)
	fmt.Println(cnt)
}





