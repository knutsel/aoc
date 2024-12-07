from utils import get_input


def print_grid():
    for r in grid:
        for c in r:
            print(c, end="")
        print("")


def walk(cur_loc, cur_dir):
    visited = set()
    visited.add(cur_loc)
    vis_with_direction = set()
    while True:
        new_loc = (cur_loc[0] + step[cur_dir][0], cur_loc[1] + step[cur_dir][1])
        if not is_inbounds(new_loc[0], new_loc[1]):
            break
        if grid[new_loc[0]][new_loc[1]] == "#":
            cur_dir = rotate[cur_dir]
        else:
            cur_loc = new_loc
            visited.add(cur_loc)
        if (cur_loc, cur_dir) in vis_with_direction:
            return visited, True
        vis_with_direction.add((cur_loc, cur_dir))
    return visited, False


def is_inbounds(y, x):
    if x < 0 or y < 0 or x >= len(grid[0]) or y >= len(grid):
        return False
    return True


data = get_input(for_example=False, day=6)
step = {"^": (-1, 0), ">": (0, 1), "<": (0, -1), "v": (1, 0)}
rotate = {"^": ">", ">": "v", "<": "^", "v": "<"}
start_loc = (-1, -1)
direction = "x"
grid = []
for y, line in enumerate(data):  # use y, x in everything --> [line-no][char-on-line]
    grid.append(list(line))
    for x, char in enumerate(line):
        if char == '^' or char == 'v' or char == '<' or char == '>':
            start_loc = y, x
            direction = char

vis, _ = walk(start_loc, direction)

p1 = len(vis)
p2 = 0
for i, v in enumerate(vis):
    grid[v[0]][v[1]] = '#'
    _, loop = walk(start_loc, direction)
    if loop:
        p2 += 1
    grid[v[0]][v[1]] = '.'

print(p1, p2)
