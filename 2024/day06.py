from aocd import get_data


def get_input(for_example):
    if for_example:
        with open("/tmp/ex") as file:
            data = file.read().splitlines()
    else:
        data = get_data(year=2024, day=6).splitlines()
    return data


def print_grid():
    for r in grid:
        for c in r:
            print(c, end="")
        print("")


def would_cause_loop(cur_loc, cur_dir):
    grid[cur_loc[0]][cur_loc[1]] = '#'
    been_before = [(cur_loc, cur_dir)]
    # print(start_loc, cur_dir, end="  ")
    new_obstacle_loc  = (cur_loc[0] + step[cur_dir][0], cur_loc[1] + step[cur_dir][1])
    cur_dir = rotate[cur_dir] # Extra turn is as if there was a blockage next
    # new_loc_if_fails  = (cur_loc[0] + step[cur_dir][0], cur_loc[1] + step[cur_dir][1])
    turn_count = 0
    while True:
        new_loc = (cur_loc[0] + step[cur_dir][0], cur_loc[1] + step[cur_dir][1])
        # print(f"    {new_loc}, {cur_dir}")
        if (new_loc, cur_dir) in been_before:
            # grid[new_obstacle_loc[0]][new_obstacle_loc[1]] = "O"
            # print(f"True! {new_obstacle_loc}")
            grid[cur_loc[0]][cur_loc[1]] = '#'
            return True
        if not is_inbounds(new_loc[0], new_loc[1]):
            # print("False -- OOB")
            return False
        if grid[new_loc[0]][new_loc[1]] == "#":
            cur_dir = rotate[cur_dir]
            turn_count += 1
        else:
            # print(f"   {cur_dir} --> {new_loc}")
            cur_loc = new_loc
        been_before.append((new_loc, cur_dir))

    print("False - END")
    return False


def walk(cur_loc, cur_dir):
    p1 = p2 = 0
    while True:
        new_loc = (cur_loc[0] + step[cur_dir][0], cur_loc[1] + step[cur_dir][1])
        if would_cause_loop(cur_loc, cur_dir):
            p2 += 1
        if not is_inbounds(new_loc[0], new_loc[1]):
            break
        if grid[new_loc[0]][new_loc[1]] == "#":
            cur_dir = rotate[cur_dir]
        else:
            if grid[new_loc[0]][new_loc[1]] == ".":  # only count when visiting the first time
                p1 += 1
            grid[new_loc[0]][new_loc[1]] = cur_dir
            cur_loc = new_loc
    return p1 + 1, p2


def is_inbounds(y, x):
    if x < 0 or y < 0 or x >= len(grid[0]) or y >= len(grid):
        return False
    return True


data = get_input(for_example=False)
step = {"^": (-1, 0), ">": (0, 1), "<": (0, -1), "v": (1, 0)}
rotate = {"^": ">", ">": "v", "<": "^", "v": "<"}
current_loc = (-1, -1)
direction = "x"
grid = []
blocked = []
for y, line in enumerate(data):
    grid.append(list(line))
    for x, char in enumerate(line):
        if char == '^' or char == 'v' or char == '<' or char == '>':
            current_loc = y, x
            direction = char
        if char == "#":
            blocked.append((y, x))

print(f"loc: {current_loc}, dir: {direction} len blocked: {len(blocked)}")

p1, p2 = walk(current_loc, direction)
print_grid()

print(p1, p2)
