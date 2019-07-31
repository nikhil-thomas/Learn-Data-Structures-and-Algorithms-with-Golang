package main

import "fmt"

type IRealObject interface {
	performAction()
}

type RealObject struct {}

func (realObject *RealObject) performAction() {
	fmt.Println("RealObject perform action")
}

type VirtualProxy struct {
	realObject *RealObject
}

func (virtualProxy *VirtualProxy) performAction() {
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = &RealObject{}
	}
	fmt.Println("Virtual proxy perform action")
	virtualProxy.realObject.performAction()
}

func main() {
	var object VirtualProxy = VirtualProxy{}
	object.performAction()
}
