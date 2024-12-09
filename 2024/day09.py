from functools import cmp_to_key

from utils import get_input


def mysorter(a, b):
    if a[1] == b[1]:
        return 0
    if a[1] > b[1]:
        return 1
    else:
        return -1


data = get_input(for_example=False, day=9)
print(data)
# files = []
# gaps = []
# for i, c in enumerate(data[0]):
#     if i % 2:
#         gaps.append(int(c))
#     else:
#         files.append((i // 2, int(c))) # id, num_blocks
#
# for g in gaps:

block_list =[]
for i, c in enumerate(data[0]):
    if i % 2:
        for _ in range(int(c)):
            block_list.append('.')
    else:
        for _ in range(int(c)):
            block_list.append(i//2)

# last = block_list.pop()
# print(block_list)
while True:
    last = block_list.pop()
    try:
        index = block_list.index('.')
    except ValueError:
        block_list.append(last)
        break
    block_list[index] = last

# print(block_list)

p1 = 0
for i, id in enumerate(block_list):
    if id == '.':
        break
    p1+= i*id

print(p1)
 # sorted_files = sorted(files, key=cmp_to_key(mysorter), reverse=False)
# sorted_gaps = sorted(gaps)
# print(files, gaps)
# print(sorted_files)
# for i, gap in enumerate(sorted_gaps):
#     g_index = sorted
