from utils import get_input, print_grid

lines = get_input(for_example=False, day=15)

move = {">": (0, 1), "<": (0, -1), "v": (1, 0), "^": (-1, 0)}
grid = []
moves = []
pos = (-1, -1)
reading_moves = False
for y, l in enumerate(lines):
    if len(l) == 0:
        reading_moves = True
    if reading_moves:
        moves.extend(list(l))
        continue
    else:
        grid.append(list(l))
        x = l.find("@")
        if x != -1:
            pos = (y, x)

# print_grid(grid)
print(moves)

for m in moves:
    # print_grid(grid)
    # print(m)
    next_pos = (pos[0] + move[m][0], pos[1] + move[m][1])
    if grid[next_pos[0]][next_pos[1]] == ".":
        # grid[next_pos[0]][next_pos[1]] = "L"
        grid[next_pos[0]][next_pos[1]] = "@"
        grid[pos[0]][pos[1]] = "."
        pos = (next_pos[0], next_pos[1])
    elif grid[next_pos[0]][next_pos[1]] == "#":
        continue
    elif grid[next_pos[0]][next_pos[1]] == "O":
        num_boxes = 1
        end_box_pos = next_pos
        while 'O' == grid[end_box_pos[0]][end_box_pos[1]]:
            end_box_pos = (end_box_pos[0] + move[m][0], end_box_pos[1] + move[m][1])
            num_boxes += 1
        if num_boxes >= 1 and grid[end_box_pos[0]][end_box_pos[1]] == ".":
            grid[next_pos[0]][next_pos[1]] = "@"
            grid[pos[0]][pos[1]] = "."
            grid[end_box_pos[0]][end_box_pos[1]] = "O"
            pos = (next_pos[0], next_pos[1])

p1 = 0
for y, _ in enumerate(grid):
    for x, _ in enumerate(grid[y]):
        if grid[y][x] == "O":
            p1 += 100 * y + x

print(p1)
