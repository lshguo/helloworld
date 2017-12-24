package main

/*
	检查一个字符串是否是回文
	hello不是回文
	oppo是回文
*/
import (
	"fmt"
)

//var str = []byte("helloworld")
var str = []byte("o")

func IsHuiWen(str []byte)bool{
	if len(str) == 1 || len(str) == 0{
		return true
	}

	L := len(str)
	if str[0] != str[L-1]{
		return false
	}else{
		subStr := str[1:L-1]
		return IsHuiWen(subStr)
	}
}
func main(){
	//len := 
	fmt.Println(IsHuiWen(str))
}
