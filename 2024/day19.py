import heapq

from utils import get_input, print_grid

def is_valid(towels, design, index):
    if index == len(design):
        return True
    for towel in towels:
        print(f"-{design[index:index+len(towel)]}- == -{towel}- (index: {index}, length: {len(towel)})")
        if design[index:index+len(towel)] == towel:
            print("YAYAY")
            if is_valid(towels, design, index + len(towel)):
                return True

    return False

data = get_input(for_example=False, day=19)
towels = data[0].split(', ')

p1 = 0
for design in data[2:]:
    print(design)
    if is_valid(towels, design, 0):
        print("YES")
        p1 += 1

print(f"Part 1: {p1}")


