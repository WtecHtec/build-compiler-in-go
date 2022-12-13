// vm/frame.go

package vm

import (
	"monkey/code"
	"monkey/object"
)

// fn指向帧引用的已编译函数，ip则表示该帧的指令指针
// vm/frame.go

// vm/frame.go

type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int
}

func NewFrame(cl *object.Closure, basePointer int) *Frame {
	f := &Frame{
		cl:          cl,
		ip:          -1,
		basePointer: basePointer,
	}

	return f
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
