package main

import "testing"

func TestEmpty(t *testing.T) {
	setCode("")
	expected := ""
	actual := validateCode("")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestUnderscore(t *testing.T) {
	setCode("4321")
	expected := "_"
	actual := validateCode("3896")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestAllUnderscore(t *testing.T) {
	setCode("2345")
	expected := "____"
	actual := validateCode("4253")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestX(t *testing.T) {
	setCode("2345")
	expected := "x"
	actual := validateCode("2719")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestXXXX(t *testing.T) {
	setCode("2345")
	expected := "xxxx"
	actual := validateCode("2345")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestXX__(t *testing.T) {
	setCode("2345")
	expected := "xx__"
	actual := validateCode("2354")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func Test__XX(t *testing.T) {
	setCode("2345")
	expected := "__xx"
	actual := validateCode("3245")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func Test__(t *testing.T) {
	setCode("2345")
	expected := "__"
	actual := validateCode("3298")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
