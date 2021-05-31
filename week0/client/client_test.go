package main

import (
	"testing"
)

func TestSetFuncField(t *testing.T) {
	helloService := &hello{

	}

	SetFuncField(helloService)

	res, err := helloService.SayHello(&Input{
		Name: "golang",
	})

	if err !=nil{
		t.FailNow()
	}


}