// code/code.go

package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte // 字节集合

type Opcode byte // 字节描述

// 操作数
const (
	OpConstant      Opcode = iota // 常量，读取栈时是读取它所在的索引
	OpAdd                         // +
	OpPop                         // 弹出
	OpSub                         // -
	OpMul                         // *
	OpDiv                         // /
	OpTrue                        // true
	OpFalse                       // false
	OpEqual                       // ==
	OpNotEqual                    // ！=
	OpGreaterThan                 // >
	OpMinus                       // -7
	OpBang                        // ！true
	OpJumpNotTruthy               // not true 跳转指令
	OpJump                        // 跳转指令

	OpNull // null

	OpGetGlobal // 声明变量
	OpSetGlobal // 获取变量

	OpArray // 数组

	OpHash // 哈希

	OpIndex // arry[i] 索引
)

type Definition struct {
	Name          string // 操作名称
	OperandWidths []int  // 操作宽度
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}}, // 常量字节码
	OpAdd:      {"OpAdd", []int{}},
	OpPop:      {"OpPop", []int{}},
	OpSub:      {"OpSub", []int{}},
	OpMul:      {"OpMul", []int{}},
	OpDiv:      {"OpDiv", []int{}},
	OpTrue:     {"OpTrue", []int{}},
	OpFalse:    {"OpFalse", []int{}},

	OpEqual:       {"OpEqual", []int{}},
	OpNotEqual:    {"OpNotEqual", []int{}},
	OpGreaterThan: {"OpGreaterThan", []int{}},

	OpMinus: {"OpMinus", []int{}},
	OpBang:  {"OpBang", []int{}},

	OpJumpNotTruthy: {"OpJumpNotTruthy", []int{2}},
	OpJump:          {"OpJump", []int{2}},

	OpNull: {"OpNull", []int{}},

	OpGetGlobal: {"OpGetGlobal", []int{2}},
	OpSetGlobal: {"OpSetGlobal", []int{2}},

	OpArray: {"OpArray", []int{2}},

	OpHash: {"OpHash", []int{2}},

	OpIndex: {"OpIndex", []int{}},
}

// 检查
func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

// 制造字节码
func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			// 256 进制 2 个字节
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}

func (ins Instructions) String() string {
	var out bytes.Buffer

	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}

		operands, read := ReadOperands(def, ins[i+1:])
		// i 操作符索引0 => 0000 ; 3 => 0003(转了进制)
		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))
		i += 1 + read
	}
	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n",
			len(operands), operandCount)
	}

	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}

	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}

// 解码
func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}

		offset += width
	}

	return operands, offset
}

func ReadUint16(ins Instructions) uint16 {
	// 2 个字节
	return binary.BigEndian.Uint16(ins)
}
