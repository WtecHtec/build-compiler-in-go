package main

import (
	"fmt"
	"strconv"
)

func VM(codes []string) {
	cLen := len(codes)
	stack := make([]string, cLen)
	cI := 0
	sI := 0
	for cI < cLen {
		curCode := codes[cI]
		switch curCode {
		case "PUSH":
			stack[sI] = codes[cI+1]
			cI++
			sI++

		case "ADD":
			right, _ := strconv.Atoi(stack[sI-1])
			sI--
			left, _ := strconv.Atoi(stack[sI-1])
			sI--
			stack[sI] = strconv.Itoa(right + left)
			sI++

		case "MINUS":
			right, _ := strconv.Atoi(stack[sI-1])
			sI--
			left, _ := strconv.Atoi(stack[sI-1])
			sI--
			stack[sI] = strconv.Itoa(left - right)
			sI++

		}
		cI++
	}
	fmt.Println(stack[sI-1])
}
func main() {
	CODES := []string{
		"PUSH", "4",
		"PUSH", "5",
		"ADD",
		"PUSH", "3",
		"MINUS",
	}
	fmt.Println("测试")
	VM(CODES)
}
