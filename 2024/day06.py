from aocd import get_data


# use y, x in everything!

def get_input(for_example):
    if for_example:
        with open("/tmp/ex") as file:
            return file.read().splitlines()
    else:
        return get_data(year=2024, day=6).splitlines()


def print_grid():
    for r in grid:
        for c in r:
            print(c, end="")
        print("")


# def would_cause_loop(cur_loc, cur_dir):
#     grid[cur_loc[0]][cur_loc[1]] = '#'
#     been_before = [(cur_loc, cur_dir)]
#     # print(start_loc, cur_dir, end="  ")
#     new_obstacle_loc = (cur_loc[0] + step[cur_dir][0], cur_loc[1] + step[cur_dir][1])
#     cur_dir = rotate[cur_dir]  # Extra turn is as if there was a blockage next
#     # new_loc_if_fails  = (cur_loc[0] + step[cur_dir][0], cur_loc[1] + step[cur_dir][1])
#     turn_count = 0
#     while True:
#         new_loc = (cur_loc[0] + step[cur_dir][0], cur_loc[1] + step[cur_dir][1])
#         # print(f"    {new_loc}, {cur_dir}")
#         if (new_loc, cur_dir) in been_before:
#             # grid[new_obstacle_loc[0]][new_obstacle_loc[1]] = "O"
#             # print(f"True! {new_obstacle_loc}")
#             grid[cur_loc[0]][cur_loc[1]] = '#'
#             return True
#         if not is_inbounds(new_loc[0], new_loc[1]):
#             # print("False -- OOB")
#             return False
#         if grid[new_loc[0]][new_loc[1]] == "#":
#             cur_dir = rotate[cur_dir]
#             turn_count += 1
#         else:
#             # print(f"   {cur_dir} --> {new_loc}")
#             cur_loc = new_loc
#         been_before.append((new_loc, cur_dir))
#
#     print("False - END")
#     return False


def walk(cur_loc, cur_dir, find_loops):
    # print(f"start: {cur_loc}, {cur_dir}", end=" ")
    visited = set()
    vis_with_direction = set()
    while True:
        new_loc = (cur_loc[0] + step[cur_dir][0], cur_loc[1] + step[cur_dir][1])
        # print(f"  new_loc: {cur_loc}, {cur_dir}", end=" ")
        if not is_inbounds(new_loc[0], new_loc[1]):
            # print("br", end=" ")
            break
        if grid[new_loc[0]][new_loc[1]] == "#":
            cur_dir = rotate[cur_dir]
            # print(f"  new_dir: {cur_dir}", end=" ")
        else:
            cur_loc = new_loc
            visited.add(cur_loc)
        if (cur_loc, cur_dir) in vis_with_direction and find_loops:
            return visited, True
        vis_with_direction.add((cur_loc, cur_dir))
    # print("ret")
    return visited, False


def is_inbounds(y, x):
    if x < 0 or y < 0 or x >= len(grid[0]) or y >= len(grid):
        return False
    return True


data = get_input(for_example=False)
step = {"^": (-1, 0), ">": (0, 1), "<": (0, -1), "v": (1, 0)}
rotate = {"^": ">", ">": "v", "<": "^", "v": "<"}
start_loc = (-1, -1)
direction = "x"
grid = []
for y, line in enumerate(data):
    grid.append(list(line))
    for x, char in enumerate(line):
        if char == '^' or char == 'v' or char == '<' or char == '>':
            start_loc = y, x
            direction = char

print(f"loc: {start_loc}, dir: {direction}")

vis, loop = walk(start_loc, direction, False)
print_grid()

p1 = len(vis)
p2 = 0
print_grid()
for i, v in enumerate(vis):
    print(f"i:{i}   {v}")
    # if v == (6,3):
    #     print_grid()
    grid[v[0]][v[1]] = '#'
    vis, loop = walk(start_loc, direction, True)
    if loop:
        p2 += 1
    grid[v[0]][v[1]] = '.'
    # grid[v[0]][v[1]] = 'X'

print()
print_grid()


print(p1, p2)
