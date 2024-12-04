from aocd import get_data

instr = get_data(year=2024, day=4).splitlines()


def check_word(y, x, word, step):
    for at_index in range(1, len(word)):
        x, y = x + step[0], y + step[1]
        if x < 0 or x > len(instr) - 1 or y < 0 or y > len(instr[0]) - 1:
            return False
        if instr[y][x] != word[at_index]:
            return False
    return True


def check_cross(y, x):  # this is somewhat verbose, but readable (at least to me ;-))
    if y + 1 > len(instr) - 1 or x + 1 > len(instr[0]) - 1 or y - 1 < 0 or x - 1 < 0:
        return False
    if not ((instr[y - 1][x - 1] == 'M' and instr[y + 1][x + 1] == 'S') or (
            instr[y - 1][x - 1] == 'S' and instr[y + 1][x + 1] == 'M')):
        return False
    if not ((instr[y - 1][x + 1] == 'M' and instr[y + 1][x - 1] == 'S') or (
            instr[y - 1][x + 1] == 'S' and instr[y + 1][x - 1] == 'M')):
        return False
    return True


p1 = p2 = 0
for y, line in enumerate(instr):
    for x, char in enumerate(line):
        if char == "X":  # pick a direction, and keep following it in check_word
            for step in ((0, 1), (0, -1), (-1, 1), (-1, 0), (-1, -1), (1, 1), (1, 0), (1, -1)):
                if check_word(y, x, "XMAS", step):
                    p1 += 1
        if char == "A":
            if check_cross(y, x):
                p2 += 1

print(p1, p2)
