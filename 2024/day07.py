from utils import get_input


def do_op(operand1, operand2, operator):
    match operator:
        case '+':
            return operand1 + operand2
        case '*':
            return operand1 * operand2
        case '|':
            return int(str(operand1) + str(operand2))


def is_valid(calculated, wanted, operation_list):
    if len(operation_list) == 0:
        if calculated == wanted:
            return True
        else:
            return False
    op1 = calculated
    op2 = operation_list[0]
    operation_list = operation_list[1:]
    for operator in '*', '+', '|':
        calculated = do_op(op1, op2, operator)
        if is_valid(calculated, wanted, operation_list):
            return True


data = get_input(for_example=False, day=7)

equations = {}
for line in data:
    result = int(line.split(':')[0])
    operands = list(map(int, line.split(':')[1].strip().split(' ')))
    equations[result] = operands  # I checked there are no dupe keys

p1 = p2 = 0
for result, operands in equations.items():
    if is_valid(operands[0], result, operands[1:]):
        p1 += result

print(p1, p2)
