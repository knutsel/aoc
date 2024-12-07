from utils import get_input


def do_op(operand1, operand2, operator):
    print(f"{operand1} {operator} {operand2}")
    match operator:
        case '+':
            return operand1 + operand2
        case '*':
            # if operand1 == 0:
            #     operand1 = 1
            return operand1 * operand2


# def is_valid(result, operands, running_total, oplist):
#     print(f"r:{result}, operands:{operands}")
#
#     for i, operand in enumerate(operands[:-1]):
#         oplist.append(operand)
#         for operator in '+', '*':
#             new_value = running_total + do_op(operand,  operands[i+1], operator)
#             if new_value < result:
#                 oplist.append(operator)
#                 # oplist.append([str(operand), str(operands[i+1]), operator])
#                 if len(operands)-i <= 2:
#                     continue
#                 if is_valid(result, operands[i+1:], new_value, oplist):
#                     continue

# total_after_op = running_total + operand * operands[i+1]
# if total_after_op < result:
# oplist.append(str(operand))
# oplist.append('*')
# oplist.append(str(operands[i+1]))
# if is_valid(result, operands[i+1:], running_total, oplist):
#     continue
# else:
#     total_after_op = running_total + op * operands[i + 1]
#     # running_total += op + operands[i+1]
#     if total_after_op > result:
#         continue
#     oplist.append(str(op))
#     oplist.append('+')
#     oplist.append(str(operands[i+1]))
#     if is_valid(result, operands[i+1:], running_total, oplist):
#         continue

# if new_value == result:
#     return True
# else:
#     return False

# g= {} # lets try with a graph
# for i, op in enumerate(operands[:-1]):
#     g[op]=['+']
#     g[op].append(operands[i+1])
#     g[op]=['*']
#     g[op].append(operands[i+1])
#
# print(g)

def is_valid(calculated, wanted, operation_list):
    if len(operation_list) == 0:
        if calculated == wanted:
            return True
        else:
            return False
    op1 = calculated
    op2 = operation_list[0]
    operation_list = operation_list[1:]
    for operator in '*', '+':
        calculated = do_op(op1, op2, operator)
        if not is_valid(calculated, wanted, operation_list):
            continue
        else: return True



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
