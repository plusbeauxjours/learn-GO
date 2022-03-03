package main

import (
	"fmt"

	"github.com/plusbeauxjours/learnGo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}

	baseWord := "hello"

	dictionary.Add(baseWord, "First")
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	
	dictionary.Update(baseWord, "Second")
	word, err2 := dictionary.Search(baseWord)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(word)
	}

	dictionary.Delete(baseWord)
	word, err3 := dictionary.Search(baseWord)

	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(word)
	}
}