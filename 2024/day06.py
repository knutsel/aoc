from aocd import get_data

data = get_data(year=2024, day=6).splitlines()

# with open("/tmp/ex") as file:
#     data = file.read().splitlines()


def is_inbounds(y, x):
    if x < 0 or y < 0 or x >= len(map[0]) or y >= len(map):
        return False
    return True


step = {"^": (-1, 0), ">": (0, 1), "<": (0, -1), "v": (1, 0)}
rotate = {"^": ">", ">": "v", "<": "^", "v": "<"}
p1 = p2 = 0
current_loc = (-1, -1)
dir = "x"
map = []
for y, line in enumerate(data):
    map.append(list(line))
    for x, char in enumerate(line):
        if char == '^' or char == 'v' or char == '<' or char == '>':
            current_loc = y, x
            dir = char

print(f"loc: {current_loc}, dir: {dir}")

p1 = 1
while True:
    new_loc = (current_loc[0] + step[dir][0], current_loc[1] + step[dir][1])
    if not is_inbounds(new_loc[0], new_loc[1]):
        break
    if map[new_loc[0]][new_loc[1]] == "#":
        dir = rotate[dir]
    else:
        if map[new_loc[0]][new_loc[1]] == ".":
            p1 += 1
        map[new_loc[0]][new_loc[1]] = dir
        current_loc = new_loc

for r in map:
    for c in r:
        print(c, end="")
    print("")

print(p1, p2)
