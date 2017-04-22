package main

import "strings"

var private_code string

func setCode(code string) {
	private_code = code
}

func validateCode(code string) string {
	codeArray := strings.Split(code, "")
	response := ""
	size := len(code)

	for i := 0; i < size; i++ {
		if strings.Index(private_code, codeArray[i]) == i {
			response += "x"
		} else if strings.Contains(private_code, codeArray[i]) {
			response += "_"
		}
	}
	return response
}
