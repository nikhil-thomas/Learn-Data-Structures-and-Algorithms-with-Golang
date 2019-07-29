package main

import "fmt"

type Iprocess interface {
	process()
}

type ProcessClass struct {}

func (pc *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

type ProcessDecorator struct {
	processInstance *ProcessClass
}

func (pd *ProcessDecorator) process() {
	if pd.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Println("ProcessDecorator process and")
		pd.processInstance.process()
	}
}

func main() {
	var process = &ProcessClass{}
	var decorator = &ProcessDecorator{}
	decorator.process()
	decorator.processInstance = process
	decorator.process()
}
