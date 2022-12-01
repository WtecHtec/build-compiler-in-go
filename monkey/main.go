package main

import (
	"encoding/binary"
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func infaCompiler() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

type Test struct {
	value string
}

func (test Test) String() string {
	return "测试"
}

const (
	OpConstant int = iota
)

func main() {
	res := make([]byte, 2)
	binary.BigEndian.PutUint16(res[0:], uint16(256))
	fmt.Print(res)
}
