import utils


def check_neighbour_is_symbol(pos):
    for x1 in range(pos[1], pos[1] + pos[2]):
        for n in (-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1):
            y = n[0] + pos[0]
            x = x1 + n[1]
            if x < 0 or y < 0 or x > len(lines[0]) - 1 or y > len(lines) - 1:
                continue
            if lines[y][x] != '.' and not lines[y][x].isnumeric():
                if lines[y][x] == '*':
                    if (y, x) in gear_loc:
                        gear_loc[(y, x)].append(pos[3])  # modifies the global
                    else:
                        gear_loc[(y, x)] = [pos[3]]
                return True
    return False


lines = utils.read_file('input03.txt')  # lines is a global var, is that bad?
gear_loc = {}  # global as well
number_locations = []  # (y, x ,len, value) tuples
for y, line in enumerate(lines):
    strlen = 0
    for x, c in enumerate(line):
        if c.isnumeric():
            strlen += 1
        else:
            if strlen > 0:
                t = (y, x - strlen, strlen, int(line[x - strlen:x]))
                number_locations.append(t)
            strlen = 0
    if strlen > 0:
        t = (y, len(line) - strlen, strlen, int(line[len(line) - strlen:]))
        number_locations.append(t)

sum_parts = 0
for num_loc in number_locations:
    if check_neighbour_is_symbol(num_loc):
        sum_parts += num_loc[3]

sum2 = 0
for gear in gear_loc:
    if len(gear_loc[gear]) == 2:  # there are no 3 geared numbers
        sum2 += gear_loc[gear][0] * gear_loc[gear][1]

print(sum_parts)
print(sum2)
