// process.stdout.write('A')

const _OPR = '>' // Pointer right  将指针向右移动
const _OPL = '<' // Pointer left   将指针向左移动
const _ADD = '+' // Add unit       增加指针处的内存单元
const _SUB = '-' // Sub unit       减少指针处的内存单元
const _OUT = '.' // Output         输出指针所在单元格所代表的字符
const _OIN = ',' // Input          输入一个字符并将其存储在指针所在的单元格中
const _OJF = '[' // Jump forward   如果指针处的单元格为零，则跳过匹配项
const _OJB = ']' // Jump back      如果指针处的单元格非零，则跳回匹配项

class opcode {
    code = 0
    jump = 0
}

function compile() {
    let program=[]
    let stack=[]
}

function execute() {
}

function run() {
}

(() => {
    console.log('run main.nodejs')
})()
