package student

import (
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name string
	Age int
	Grade int
}
func NewStudent(l string) *Student{
	s:= strings.Split(l," ")
	if len(s) !=3{
		fmt.Println("введены не корректные данные")
		return nil
	}
	p := Student{}
	p.Name = s[0]
	p.Age,_ = strconv.Atoi(s[1])
	p.Grade ,_ = strconv.Atoi(s[2])
	return &p
}