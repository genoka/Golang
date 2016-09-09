package main

import "C"

func bricolage(argc int,argv string) int{
 sum := func(a, b int) int { return a+b } (300, 700)
 return sum
}
func Random() int {
    return int( C.random())
}
func main(){
sum := func(a, b int) int { return a+b } (3, 4)
println(sum)
println("a la ligne")
println( bricolage(3,"jjjj") )
print( Random() )
}