package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string)(int, string){
	defer fmt.Println("I'm done")
	return len(name), strings.ToUpper(name)
}

func lenAndUpperV2(name string)(length int, uppsercase string){
	defer fmt.Println("I'm done")
	length = len(name)
	uppsercase = strings.ToUpper(name)
	return 
}

func multiply(a int,b int) int{
	return a*b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main(){
	totlalLength, upperName := lenAndUpper("hello")
	totlalLengthV2, upperNameV2 := lenAndUpperV2("hello")
	name := "mj"
	name = "hannah"
	fmt.Println(name)
	fmt.Println(multiply(4,5))
	fmt.Println(totlalLength,upperName)
	fmt.Println(totlalLengthV2,upperNameV2)
	repeatMe("tomato", "apple", "orange", "banana")
} 