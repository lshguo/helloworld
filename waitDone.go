package main

import (
	"fmt"
	"time"
	"sync"
)

type Foo struct{
	x int
	m sync.Mutex
	w sync.WaitGroup
}
func (f *Foo)Run(){
	f.m.Lock()
	f.x = f.x + 1
	time.Sleep(time.Second)
	f.m.Unlock()
	f.w.Done()
}

func main(){
	f := new(Foo)
	cnt := 3
	for i := 0; i < cnt; i++{
		f.w.Add(1)
		go f.Run()
	}
	f.w.Wait()
	fmt.Println(f.x)
}


