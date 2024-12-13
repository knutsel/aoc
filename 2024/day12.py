from utils import get_input

data = get_input(for_example=False, day=12)


def is_inbounds(y, x):
    if x < 0 or y < 0 or x >= col_num or y >= row_num:
        return False
    return True


def define_region(ptype, row, col, region, grid):  # is this bfs or dfs?
    if grid[row][col] != ptype:
        return
    region.add((row, col))
    for (r, c) in (row, col + 1), (row, col - 1), (row - 1, col), (row + 1, col):
        if is_inbounds(r, c) and grid[r][c] == ptype and (r, c) not in region:
            region.add((r, c))
            define_region(ptype, r, c, region, grid)

    return


def perimeter(region):
    perimeter = 0
    for (row, col) in region:
        for (r, c) in (row, col + 1), (row, col - 1), (row - 1, col), (row + 1, col):
            if (r, c) not in region:
                perimeter += 1

    return perimeter


row_num = len(data)
col_num = len(data[0])
total_plots = row_num * col_num
visited = set()
total_price = 0
for row in range(row_num):
    for col in range(col_num):
        if (row, col) in visited:
            continue
        region = set()
        define_region(data[row][col], row, col, region, data)
        total_price += perimeter(region) * len(region)
        visited = visited | region

print(total_price)
