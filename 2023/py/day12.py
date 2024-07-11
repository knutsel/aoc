import re

import utils

def match_pattern(inputline, pattern):
    restring=''
    for p in pattern:
        restring+='#{'+str(p)+'}\.+'
    if re.match(restring,inputline):
        return True
    return False
    # hashes_to_find = pattern.pop(0)
    # for c in inputline:
    #     # try:
    #     # except IndexError:
    #     #     return True
    #     if c == '#' or c == '?':
    #         hashes_to_find-=1
    #     else:
    #         if hashes_to_find != 0:
    #             return False
    #         # hashes_to_find=pattern.pop(0)

def count_possible(inputline, pattern):
    print(f'in: {inputline}')
    try:
        q=inputline.index('?')
    except ValueError:
        print(f'out: {inputline}')
        return 1

    # if inputline.find('?') == -1:
    #     print(f'out: {pattern}')
    #     return 1

    for i, q in enumerate(inputline):
        if q == '?':
            newline = inputline[:]
            newline[i] = '#'
            if match_pattern(newline, pattern):
                count_possible(newline, pattern)
            newline[i] = '.'
            if match_pattern(newline, pattern):
                count_possible(newline, pattern)

    # print(f'ELSE: {inputline}')

    # print(f'OUT: {inputline}')


lines: list[str] = utils.read_file('example.txt')
for line in lines:
    p1, p2 = line.split(' ')
    # qlocs = p1.findall('?')
    # qlocs = list(re.finditer(p1, '\\?'))
    print(f'<>>> p1: {p1}, p2: {p2}')
    possible = count_possible(list(p1), list(map(int, p2.split(','))))
