// main.rs

use std::io;
use std::io::Write;

const _OPR: char = '>';
const _OPL: char = '<';
const _ADD: char = '+';
const _SUB: char = '-';
const _OUT: char = '.';
const _OIN: char = ',';
const _OJF: char = '[';
const _OJB: char = ']';

struct Opcode {
    code: char,
    jump: usize,
}

fn compile(input: String) -> Vec<Opcode> {
    let mut program: Vec<Opcode> = vec![];
    let mut stack: Vec<usize> = vec![];
    let mut jump: usize;
    let mut pc: usize = 0;
    for char in input.chars() {
        match char {
            _OPR => program.push(Opcode { code: _OPR, jump: 0 }),
            _OPL => program.push(Opcode { code: _OPL, jump: 0 }),
            _ADD => program.push(Opcode { code: _ADD, jump: 0 }),
            _SUB => program.push(Opcode { code: _SUB, jump: 0 }),
            _OUT => program.push(Opcode { code: _OUT, jump: 0 }),
            _OIN => program.push(Opcode { code: _OIN, jump: 0 }),
            _OJF => {
                program.push(Opcode { code: _OJF, jump: 0 });
                stack.push(pc);
            }
            _OJB => {
                if stack.len() == 0 {
                    return program;
                }
                let index = stack.len() - 1;
                jump = stack[index];
                stack = Vec::from(&stack[..index]);
                program.push(Opcode { code: _OJB, jump: jump });
                program[jump].jump = pc;
            }
            _ => pc += 1
        }
        pc -= 1
    }
    if stack.len() != 0 {
        return program;
    }
    return program;
}

fn execute(program: Vec<Opcode>) {
    // let reader = bufio.NewReader(os.Stdin);
    let mut data: Vec<i32> = vec![];
    let mut data_ptr = 0;
    for mut i in 0..program.len() {
        match program[i].code {
            _OPR => data_ptr += 1,
            _OPL => data_ptr -= 1,
            _ADD => data[data_ptr] += 1,
            _SUB => data[data_ptr] -= 1,
            _OUT => print!("{}", data[data_ptr]),
            _OIN => {}
            _OJF => {
                if data[data_ptr] == 0 {
                    i = program[i].jump;
                }
            }
            _OJB => {
                if data[data_ptr] > 0 {
                    i = program[i].jump;
                }
            }
            _ => panic!("compileError")
        }
    }
}

fn main() {
    let data = cat("../testdata/hello_world.bf");
    let program = compile(data);
    execute(program);
}

fn cat(fp: &str) -> String {
    let bytes = std::fs::read(fp).unwrap();
    String::from_utf8(bytes).unwrap()
}
