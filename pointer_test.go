package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

type PointerObj struct {
	Name string
	Age  int
}

type MyObject struct {
	Filed int
}

func TestObj(t *testing.T) {
	obj := PointerObj{Name: "tester", Age: 100}
	ptr := unsafe.Pointer(&obj)

	object := objc.ObjectFrom(ptr)

	tobj := (*MyObject)(object.Ptr())
	fmt.Println("myobj", tobj.Filed)

	fmt.Println(isObj(ptr))
}

func isObj(ptr unsafe.Pointer) bool {
	typ := reflect.TypeOf(ptr)
	// 判断对象类型是否为 MyObject
	return typ == reflect.TypeOf(&PointerObj{})
}
