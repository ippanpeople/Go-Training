package main

import "fmt"

type book string

func (b book) printTitle() {
	fmt.Println(b)
}

type laptopSize float64

func (l laptopSize) getSizeOfLaptop() laptopSize {
	return l
}

func main() {
	c := book("123")
	c.printTitle()
	i := laptopSize(3.14)
	i.getSizeOfLaptop()
}
