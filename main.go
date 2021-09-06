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
func main() {
	argsWithProg := os.Args
	i := len(os.Args[1:])
	fmt.Println(argsWithProg)
	fmt.Println(i)
	if i == 1 {
		file, err := os.Open(os.Args[i])
		if err != nil {
			panic(err)
		}
		defer file.Close()
		resultBytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(resultBytes))
	} else if i == 2 {
		file, err := os.Open(os.Args[i-1])
		if err != nil {
			log.Fatal()
		}
		defer file.Close()
		file2, err := os.Open(os.Args[i])
		if err != nil {
			log.Fatal()
		}
		defer file2.Close()
		resultBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal()
		}
		resultBytes2, err := ioutil.ReadAll(file2)
		if err != nil {
			log.Fatal()
		}
		splits := []string{
			string(resultBytes),
			string(resultBytes2),
		}
		splits2 := strings.Join(splits, " ")
		fmt.Println(splits2)
	} else if i == 3 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file2, err := os.Open(os.Args[2])
		if err != nil {
			panic(err)
		}
		defer file2.Close()
		resultBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal()
		}
		resultBytes2, err := ioutil.ReadAll(file2)
		if err != nil {
			log.Fatal()
		}
		splits := []string{
			string(resultBytes),
			string(resultBytes2),
		}
		splits2 := strings.Join(splits, " ")
		fmt.Println(splits2)
		file3, err := os.Create("thirdFile.txt")
		if err != nil {
			log.Fatal()
		}
		defer file3.Close()
		file3.WriteString(splits2)
		file.Sync()
	}
}
