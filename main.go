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

func superAdd(numbers ...int) int{
	for index, number:=range numbers{
		fmt.Println("index",index,"number",number)
	}
	for _, number:=range numbers{
		fmt.Println("number", number)
	}
	total :=0
	for _, number:=range numbers{
		total+=number
	}
	return total
}

func canIDrink (age int)bool{
	if koreanAge :=age+2; koreanAge<18{  
		return false
	}
	return true
}

func canIDrinkSwitch (age int)bool{
	switch koreanAge:= age+2; koreanAge{
	case 18:
		return false
	case 10:
		return true
	case 50:
		return true
	}
	return false
}


func main(){
	name := "mj"
	name = "hannah"
	fmt.Println(name)

	totlalLength, upperName := lenAndUpper("hello")
	fmt.Println(totlalLength,upperName)

	totlalLengthV2, upperNameV2 := lenAndUpperV2("hello")
	fmt.Println(totlalLengthV2,upperNameV2)
	
	fmt.Println(multiply(4,5))

	repeatMe("tomato", "apple", "orange", "banana")

	result :=superAdd(1,2,3,4,5)
	fmt.Println("result", result)

	fmt.Println(canIDrink(17))
	fmt.Println(canIDrinkSwitch(17))
} 