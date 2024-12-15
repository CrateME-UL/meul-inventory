package main

import (
	"encoding/json"
	"fmt"
)

type foo struct {
	Message    string
	Ports      []int
	ServerName string
}

func newFoo() (*foo, error) {
	return &foo{
		Message:    "foo loves bar",
		Ports:      []int{80},
		ServerName: "Foo",
	}, nil
}

func main() {
	res, err := newFoo()

	if err != nil {
		panic("trash")
	}

	out, err := json.Marshal(res)
	if err != nil {
		panic("garabage")
	}
	fmt.Printf("out = %+v\n", out)
	fmt.Printf("res = %+v\n", res)
}
