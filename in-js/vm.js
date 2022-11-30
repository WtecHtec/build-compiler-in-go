const CODE = [
  'PUSH', 5,
  'PUSH', 4,
  'ADD',
  'PUSH', 3,
  'MINUS'
]

function VM(codes) {
  const stack = []
  let cI = 0
  let sI = 0
  const len = codes.length
  while(cI < len) {
    const curCode = codes[cI]
    console.log('curCode==', curCode)
    switch(curCode) {
      case 'PUSH':
        // 入栈
        stack[sI] = codes[cI + 1]
        cI++;
        cI++;
        sI++;
        break
      case 'ADD':
        // 出栈
        stack[sI] = stack[sI - 1] + stack[sI - 2]
        cI++;
        sI++;
        break
      case 'MINUS':
        // 出栈
        stack[sI] = stack[sI - 2] - stack[sI - 1]
        cI++;
        sI++;
        break
    }
  }
  console.log(stack)
}

VM(CODE)