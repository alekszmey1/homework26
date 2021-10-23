package main

import (
"bufio"
"fmt"
	"homework.28/storage"
	"homework.28/student"
	"os"
)
func main() {
	university := storage.NewUniversity()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		l := student.NewStudent(line)
		fmt.Println(l)
		university.Put(l)
	}
	for _, v := range university.GetAll() {
		fmt.Println(v )
	}
}