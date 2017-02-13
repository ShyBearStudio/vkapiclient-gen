package main

import (
	//"fmt"
	"io/ioutil"
)

func main() {
	var b []byte
	s := IntegerTypePrimitive{1, 2}
	op1 := ObjectPropertyPrimitive{"PropertyName1", "Description1", true, &s}
	op2 := ObjectPropertyPrimitive{"PropertyName2", "Description2", true, &s}
	op3 := ObjectPropertyPrimitive{"PropertyName3", "Description3", true, &s}
	data := ObjectPrimitive{
		"SomeObjectName",
		[]*ObjectPropertyPrimitive{&op1, &op2, &op3}}
	b = CreateObject(data)
	ioutil.WriteFile("result/objects.go", b, 0666)
}
