from utils import get_input, print_grid

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
            if (r,c) not in region:
                perimeter+=1

    return perimeter
    # x = set()
    # y = set()
    # for (r, c) in region:
    #     x.add(c)
    #     y.add(r)
    # return 2* (len(y)+ len(x))

row_num = len(data)
col_num = len(data[0])
total_plots = row_num * col_num
# print(f" r:{row_num}, c:{col_num} d:\n{data} ")
print_grid(data)
visited = set()
start = (0, 0)
all_regions = {}
plant_type = data[0][0]
total_price = 0
# while len(visited) < total_plots:
for row in range(row_num):
    for col in range(col_num):
        if (row, col) in visited:
            continue
        region = set()
        define_region(data[row][col], row, col, region, data)
        print(f"{data[row][col]}: {region} {len(region)} -> {perimeter(region)} {len(visited)}")
        total_price += perimeter(region)*len(region)
        visited = visited | region

#             if plant_type == data[row][col]:
#                 region.add((row, col))


define_region(data[0][0], 0, 0, region, data)
print(total_price)
