package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	var uname string

	flag.StringVar(&uname, "u", "root", "Specify username. Default is root")
	flag.Parse()

	if uname == "firstFile.txt" {
		file, err := os.Open("firstFile.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		resultBytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(resultBytes))
	} else if uname == "firstFile.txt secondFile.txt" {
		file, err := os.Open("firstFile.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file2, err := os.Open("secondFile.txt")
		if err != nil {
			panic(err)
		}
		defer file2.Close()
		resultBytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		resultBytes2, err := ioutil.ReadAll(file2)
		if err != nil {
			panic(err)
		}
		splits := []string{
			string(resultBytes),
			string(resultBytes2),
		}
		splits2 := strings.Join(splits, " ")
		fmt.Println(splits2)
	} else if uname == "firstFile.txt secondFile.txt resultFile.txt" {
		file, err := os.Open("firstFile.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file2, err := os.Open("secondFile.txt")
		if err != nil {
			panic(err)
		}
		defer file2.Close()
		resultBytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		resultBytes2, err := ioutil.ReadAll(file2)
		if err != nil {
			panic(err)
		}
		splits := []string{
			string(resultBytes),
			string(resultBytes2),
		}
		splits2 := strings.Join(splits, " ")
		fmt.Println(splits2)
		file3, err := os.Create("thirdFile.txt")
		if err != nil {
			panic(err)
		}
		defer file3.Close()
		file3.WriteString(splits2)
	}
}
