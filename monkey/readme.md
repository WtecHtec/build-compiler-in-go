## Uint16 存储 2 个字节
# 调用栈
代码： 1 + 2  => [1 , 2, +] => [3]
索引移动，遇到常量入栈，遇到操作符出栈，操作结果入栈
# 文件夹 code
定义字节码以及其映射关系
## 生成指令 Make
生成字节码，格式： 
```
 [op, opWidth..]
常量  {
  op, //操作类型
  opWidth, // 字节长度
} => 2 => [OP, 2] => [OP, 0, 2] // 256进制

```
## 解码 ReadOperands
常量=> [OP, 0, 2] => 2
切片 1 开始转换数据
# 文件夹 compiler
生成字节码数组类型结构
编译器,将ast抽象树编译成字节码的数据格式
instructions code.Instructions // 生成的字节码
constants    []object.Object   // 数据池
数据将会在 instructions 中以索引的形式存储, 取数时根据索引从数据池中取数
# 文件夹 vm
虚拟机,遍历 compiler 生成的字节码数组，实现调用栈，执行代码。
```
 type VM struct {
	constants    []object.Object // 数据池
	instructions code.Instructions // 生成的字节码

	stack []object.Object //  存储数据类型的栈
	sp    int             // 始终指向栈中的下一个空闲槽。栈顶的值是stack[sp-1]
}
```
通过sp 指针索引，实现入栈 出栈。遍历 instructions 字节码数组，通过判断不同的操作符，调用stack出入栈，实现代码运行

# 字节码转换格式
## 常量
[opt, value]
value 存放在  constants 数据池中，
instructions 存放的是constants 数据池中value的索引
## 算式
[value, value, opt]
## 比较
[value, value, opt]
## if 条件语句
[con , jump_not, ex jump, ex]
新增 跳转指令, 指定跳转位置的索引。
[ 4 , > , 5, { jump_not: 6 }, 1 , + , 1, { jump: 12}, 1 , +, 2,]