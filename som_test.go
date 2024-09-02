package main

import "testing"

func Test_map(t *testing.T) {
	m := make(map[int]int, 2)
	t.Logf("m: %p", m)
	test := map[int]int{
		1: 2,
	}
	m = test
	t.Logf("m: %p", m)
	test[1] = 3
	t.Logf("m: %d", m[1])
	t.Log(len(m))
}

type baseInterface interface {
	SetName(name string)
}

type baseObj struct {
	Name string
}

func (b *baseObj) SetName(name string) {
	b.Name = name
}

type testObj struct {
	baseObj
}

func NewTestObj() baseInterface {
	return &testObj{}
}

func Test_obj(t *testing.T) {
	test := NewTestObj()
	test.SetName("abc")
	t.Log(test.Name)
}
