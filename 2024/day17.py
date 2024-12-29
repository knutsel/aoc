def combo(operand):
    if operand < 4:
        return operand
    if operand == 4:
        return A
    if operand == 5:
        return B
    if operand == 6:
        return C
    print("FAULT@")
    return -1


def adv(operand):
    global A, B, C, PC
    if debug: print(f"A:{A} B:{B} C:{C} PC:{PC} -adv({operand})")
    A = int(A / pow(2, combo(operand)))


def bxl(operand):
    global A, B, C, PC
    if debug: print(f"A:{A} B:{B} C:{C} PC:{PC} bxl({operand})")
    B = B | operand


def bst(operand):
    global A, B, C, PC
    if debug: print(f"A:{A} B:{B} C:{C} PC:{PC} bst({operand})")
    B = combo(operand) % 8


def jnz(operand):
    global A, B, C, PC
    if debug: print(f"A:{A} B:{B} C:{C} PC:{PC} jnz({operand})")
    if A == 0:
        return
    PC = operand - 2


def bxc(operand):
    global A, B, C, PC
    if debug: print(f"A:{A} B:{B} C:{C} PC:{PC} bxc({operand})")
    C = B | C


def out(operand):
    global A, B, C, PC
    output = combo(operand) % 8
    if debug: print(f"A:{A} B:{B} C:{C} PC:{PC} out({operand})")
    print(f"{output}", end=",")


def bdv(operand):
    global A, B, C, PC
    if debug: print(f"A:{A} B:{B} C:{C} PC:{PC} bdv({operand})")
    B = int(A / pow(2, combo(operand)))


def cdv(operand):
    global A, B, C, PC
    if debug: print(f"A:{A} B:{B} C:{C} PC:{PC} cdv({operand})")
    C = int(A / pow(2, combo(operand)))



instruction_set = ['adv', 'bxl', 'bst', 'jnz', 'bxc', 'out', 'bdv', 'cdv']
# PC and registers are global
# global A, B, C, PC
PC = 0
debug = False

A = 38610541
B = 0
C = 0
program = [2,4,1,1,7,5,1,5,4,3,5,5,0,3,3,0]
# A = 729
# B = 0
# C = 0
# program = [0, 1, 5, 4, 3, 0]

# print(f"part1: {program} ")

while PC < len(program):
    eval(instruction_set[program[PC]])(program[PC + 1])
    PC += 2
