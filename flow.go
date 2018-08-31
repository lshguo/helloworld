package main

import (
	"fmt"
<<<<<<< HEAD
	"time"
)

func main(){
	a := 1
	HERE: for{
		switch a {
		case 1, -1:
			fmt.Println(a)
			a = a + 1
		case 2, -2:
			fmt.Println(a)
			a = a + 1
		default:
			fmt.Println("...")
			if a > 0{
				a = -1 * a
				a = a + 1
			}else{
				break HERE
			}
		}
	}
}

func func1(){
	fmt.Println("hello")
	for{
		time.Sleep(time.Second)
		fmt.Println("hello")
	}
}

func func2(){
        fmt.Println("    world")
        for{
                time.Sleep(time.Second * 5)
                fmt.Println("    world")
		panic("oops!")
        }
=======
)

func main(){
	fmt.Println("hello world")
>>>>>>> 9090275671e8893fc5eeb9d2529272f8948471ed
}
