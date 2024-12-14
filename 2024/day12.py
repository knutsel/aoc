from utils import get_input, print_grid

grid = get_input(for_example=False, day=12)


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


def perimeter(region):
    perimeter = numsides = 0
    for (row, col) in region:
        # Perimeter
        perimeter += (row - 1, col) not in region
        perimeter += (row + 1, col) not in region
        perimeter += (row, col - 1) not in region
        perimeter += (row, col + 1) not in region
        # Outer corners
        numsides += (row - 1, col) not in region and (row, col - 1) not in region
        numsides += (row + 1, col) not in region and (row, col - 1) not in region
        numsides += (row - 1, col) not in region and (row, col + 1) not in region
        numsides += (row + 1, col) not in region and (row, col + 1) not in region
        # Inner corners
        numsides += (row - 1, col) in region and (row, col - 1) in region and (row - 1, col - 1) not in region
        numsides += (row + 1, col) in region and (row, col - 1) in region and (row + 1, col - 1) not in region
        numsides += (row - 1, col) in region and (row, col + 1) in region and (row - 1, col + 1) not in region
        numsides += (row + 1, col) in region and (row, col + 1) in region and (row + 1, col + 1) not in region
    return perimeter, numsides


row_num = len(grid)
col_num = len(grid[0])
total_plots = row_num * col_num
visited = set()
total_price = 0
total_price_with_discount = 0
for row in range(row_num):
    for col in range(col_num):
        if (row, col) in visited:
            continue
        region = set()
        define_region(grid[row][col], row, col, region, grid)
        peri, nsides = perimeter(region)
        total_price += peri * len(region)
        total_price_with_discount += nsides * len(region)
        visited = visited | region
print(total_price, total_price_with_discount)
