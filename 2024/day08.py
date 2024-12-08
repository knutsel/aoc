from utils import get_input


def print_grid():
    print(f"lines:{len(grid)}, cols:{len(grid[0])}")
    for r in grid:
        for c in r:
            print(c, end="")
        print("")


def is_inbounds(y, x):
    if x < 0 or y < 0 or x > len(grid[0])-1 or y > len(grid)-1:
        return False
    return True

def mirror(p1, p2):
    mirrored_x = 2 * p2[0] - p1[0]
    mirrored_y = 2 * p2[1] - p1[1]
    return (mirrored_y, mirrored_x)

def mirror_point(point_to_mirror, base_point):
    # Calculate the mirrored point
    mirrored_x = 2 * base_point[0] - point_to_mirror[0]
    mirrored_y = 2 * base_point[1] - point_to_mirror[1]
    return (mirrored_y, mirrored_x)
def get_resonator(p1, p2): # only return one, mirror it across the other
    mirrored_x, mirrored_y = mirror_point(p1, p2)
    if is_inbounds(mirrored_y, mirrored_x):
        return mirrored_y, mirrored_x

    # y2 >=
    # diffx = abs(p1[1] - p2[1])
    # r1 = r2 = (-1, -1)
    # if p2[1] < p1[1]:
    #     r1 = (p1[0] + diffy, p1[1] + diffx)
    #     r2 = (p2[0] - diffy, p2[1] - diffx)
    # else:
    #     r1 = (p2[0] + diffy, p2[1] + diffx)
    #     r2 = (p1[0] - diffy, p1[1] - diffx)
    # if p1[1] > p2[1]:
    #     r1 = (p1[0] - diffy, p2[1] + diffx)
    #     r2 = (p2[0] + diffy, p2[1] - diffx)
    # else:
    #     r1 = (p1[0] - diffy, p1[1] - diffx)
    #     r2 = (p2[0] + diffy, p1[1] + diffx)

    # retval = []
    # if is_inbounds(r1[0], r1[1]):
    #     retval.append((r1[0], r1[1]))
    #
    # if is_inbounds(r2[0], r2[1]):
    #     retval.append((r2[0], r2[1]))
    # print(f" ---> {p1},{p2} -> {retval}")
    #
    # return retval


data = get_input(for_example=False, day=8)
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
            reflection = mirror_point(a1,a2)
            if is_inbounds(reflection[0], reflection[1]):
                rset.add(tuple(reflection))
                # resonators.append((reflection[0], reflection[1]))
            # if r := get_resonator(a1, a2):
            #     resonators.append(r)
            # for r in get_resonator(a1, a2):
            #     resonators.append(r)
            # diff = (a1[0]-a2[0], a1[1]-a2[1])
            # if is_inbounds(diff[0], diff[1]):
            #     grid[diff[0]][diff[1]] = "#"
            # print(a1, a2, diff)
# rset = set()
# for r in resonators:
#     rset.add(tuple(r))
#     print(r)
#     grid[r[0]][r[1]] = "#"
p1 = len(rset)
for r in rset:
    grid[r[0]][r[1]] = "#"
print_grid()
print(p1, p2)
