package main

import (
	"fmt"
	"testing"

	"github.com/gz4z2b/algorithm/obj"
)

type testB struct {
	*obj.TestObj
	c string
}

func testObjA(o obj.TestObj) {
	fmt.Println("obj")
}

func TestStruct(t *testing.T) {
	test := testB{
		TestObj: obj.NewTestObj("a", "b"),
		c:       "111",
	}
	testObjA(*test.TestObj)
	t.Log(test.EchoA())
}
