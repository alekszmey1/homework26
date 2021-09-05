package main

import (
	"fmt"
	"os"
)

func main()  {
	text := "Go (Golang) — это многопоточный язык программирования, который был разработан компанией Google."
	text2 := "Широкой публике был предоставлен в 2009 году, и многие специалисты уверены, что будущее за ним."
	file, err := os.Create("firstFile.txt")
	if err != nil {
		fmt.Println("не смогли создать файл", err)
		return
	}
	defer file.Close()
	file.WriteString(text)
	fmt.Println(file.Name())
	file2, err := os.Create("secondFile.txt")
	if err != nil {
		fmt.Println("не смогли создать файл", err)
		return
	}
	defer file2.Close()
	file2.WriteString(text2)
	fmt.Println(file2.Name())
}