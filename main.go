package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*Программа должна получать на вход имена двух файлов, необходимо  конкатенировать их содержимое,
используя strings.Join.
Что оценивается
При получении одного файла на входе программа должна печатать его содержимое на экран.
При получении двух файлов на входе программа соединяет их и печатает содержимое обоих файлов на экран.
Если программа запущена командой go run firstFile.txt secondFile.txt resultFile.txt,
то она должна написать два соединённых файла в результирующий.*/
func treatment(a string)[]byte{
	file, err := os.Open(a)
	if err != nil {
		log.Fatal("ошибка открытия файла")
	}
	defer file.Close()
	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("ошибка чтения файла")
	}
	return resultBytes
	}

func concatenation(a[]byte, b[]byte)string{
	splits := []string{
		string(treatment(os.Args[1])),
		string(treatment(os.Args[2])),
	}
	splits2 := strings.Join(splits, " ")
	fmt.Println(splits2)
	return splits2
}

func main() {
	i := len(os.Args[1:])
	if i == 1 {
		fmt.Println(string(treatment(os.Args[1])))
	} else if i == 2 {
		concatenation(treatment(os.Args[1]),treatment(os.Args[2]))
	}else if i == 3 {
		f:= os.Args[i]
		file, err := os.Create(f)
		if err != nil {
			log.Fatal("ошибка открытия файла")
		}
		defer file.Close()
		file.WriteString(concatenation(treatment(os.Args[1]),treatment(os.Args[2])))
		file.Sync()
	}
}
