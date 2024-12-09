from utils import get_input


def is_inbounds(y, x):
    if x < 0 or y < 0 or x >= len(grid[0]) or y >= len(grid):
        return False
    return True


def reflect(point1, point2):
    mirrored_y = 2 * point2[0] - point1[0]
    mirrored_x = 2 * point2[1] - point1[1]
    return mirrored_y, mirrored_x


def nline(point1: tuple[int, int], point2: tuple[int, int]) -> set:
    lset = set()
    y1, x1 = point1
    y2, x2 = point2
    lset.add((y2, x2))
    newx = x2 + (x2 - x1)
    newy = y2 + (y2 - y1)
    while is_inbounds(newy, newx):
        lset.add((newy, newx))
        newx += (x2 - x1)
        newy += (y2 - y1)
    return lset


data = get_input(for_example=False, day=8)
grid = []
antennas = {}
for y, line in enumerate(data):  # use y, x in everything --> [line-no][char-on-line]
    grid.append(list(line))
    for x, char in enumerate(line):
        if char == ".":
            continue
        if char in antennas.keys():
            antennas[char].append((y, x))
        else:
            antennas[char] = [(y, x)]

reflection_set = set()
line_set = set()
for a in antennas:
    for a1 in antennas[a]:
        for a2 in antennas[a]:
            if a1 == a2:
                continue
            line_set = line_set.union(nline(a1, a2), nline(a1, a2))
            reflection = reflect(a1, a2)
            if is_inbounds(reflection[0], reflection[1]):
                reflection_set.add(tuple(reflection))

print(len(reflection_set), len(line_set))
