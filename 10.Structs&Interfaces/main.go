package main

import (
	"fmt"

)

type gasEngine struct {
	mpg uint8
	gallons uint8
	owner
}

type owner struct {
	name string
}

type eletricEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func (e eletricEngine) milesLeft() uint8 {
	return e.kwh*e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e gasEngine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 30, owner{"Ayush"}}

	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)

	canMakeIt(myEngine, 50)
} 