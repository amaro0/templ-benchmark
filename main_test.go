package main

import (
	"fmt"
	"testing"
)

func TestHelloTempl(t *testing.T) {
	res := helloTempl()
	if res != "<div>Hello, world</div>" {
		t.Errorf("helloTempl() = %s; want <div>Hello, world</div>", res)
	}
}

func TestHelloNative(t *testing.T) {
	res := helloNative()
	if res != "<div>Hello, world</div>" {
		t.Errorf("helloTempl() = %s; want <div>Hello, world</div>", res)
	}
}

func TestNativeTable(t *testing.T) {
	res := tableNative()
	fmt.Print(res)
}

func TestTemplTable(t *testing.T) {
	res := tableTempl()
	fmt.Print(res)
}

func TestNativeLayout(t *testing.T) {
	res := layoutNative()
	fmt.Print(res)
}

func TestTemplLayout(t *testing.T) {
	res := layoutTempl()
	fmt.Print(res)
}
