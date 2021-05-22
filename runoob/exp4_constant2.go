package main

import "unsafe"
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)
func constant()(string,int,int){
	return "abcd",3,3
}
func main(){
	x,y,z := constant()
	println(a, b, c)
	println(x, y, z)
}