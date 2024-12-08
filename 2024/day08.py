from utils import get_input


def print_grid():
    print(f"lines:{len(grid)}, cols:{len(grid[0])}")
    for r in grid:
        for c in r:
            print(c, end="")
        print("")


def is_inbounds(y, x):
    if x < 0 or y < 0 or x > len(grid[0]) - 1 or y > len(grid) - 1:
        return False
    return True


def reflect(point_to_mirror, base_point):
    mirrored_y = 2 * base_point[0] - point_to_mirror[0]
    mirrored_x = 2 * base_point[1] - point_to_mirror[1]
    return mirrored_y, mirrored_x


data = get_input(for_example=True, day=8)
grid = []
antennas = {}
p1 = p2 = 0
for y, line in enumerate(data):  # use y, x in everything --> [line-no][char-on-line]
    grid.append(list(line))
    for x, char in enumerate(line):
        if char == ".":
            continue
        if char in antennas.keys():
            antennas[char].append((y, x))
        else:
            antennas[char] = [(y, x)]

resonators = []
print_grid()
rset = set()
for a in antennas:
    print(a, antennas[a])
    for a1 in antennas[a]:
        for a2 in antennas[a]:
            if a1 == a2:
                continue
            reflection = reflect(a1, a2)
            if is_inbounds(reflection[0], reflection[1]):
                rset.add(tuple(reflection))

p1 = len(rset)
for r in rset:
    grid[r[0]][r[1]] = "#"
print_grid()
print(p1, p2)
