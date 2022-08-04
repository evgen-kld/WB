package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	test := "a4bc2d5e"
	result := "aaaabccddddde"
	mas, _ := getList(strings.Split(test, ""))
	if strings.Join(mas, "") != result {
		t.Errorf("ps.Unpack() == %q, want %q", test, result)
	}
}

func Test2(t *testing.T) {
	test := `qwe\\5`
	result := `qwe\\\\\`
	mas, _ := getList(strings.Split(test, ""))
	if strings.Join(mas, "") != result {
		t.Errorf("ps.Unpack() == %q, want %q", test, result)
	}
}

func Test3(t *testing.T) {
	test := `qwe\\4\\5`
	result := "qwe45"
	mas, _ := getList(strings.Split(test, ""))
	if strings.Join(mas, "") != result {
		t.Errorf("ps.Unpack() == %q, want %q", test, result)
	}
}

func Test4(t *testing.T) {
	test := "abcd"
	result := "abcd"
	mas, _ := getList(strings.Split(test, ""))
	if strings.Join(mas, "") != result {
		t.Errorf("ps.Unpack() == %q, want %q", test, result)
	}
}
