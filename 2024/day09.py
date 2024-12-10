from utils import get_input


def bchecksum(blist):
    sum = 0
    for i, id in enumerate(block_list):
        if id == '.':
            continue
        sum += i * id
    return sum


def mysorter(a, b):
    if a[1] == b[1]:
        return 0
    if a[1] > b[1]:
        return 1
    else:
        return -1


data = get_input(for_example=False, day=9)
# print(data)
files = []
gaps = []
# for i, c in enumerate(data[0]):
#     if i % 2:
#         gaps.append(int(c))
#     else:
#         files.append((i // 2, int(c))) # id, num_blocks
#
# for g in gaps:

block_list = []
at_block = 0
for i, c in enumerate(data[0]):
    size = int(c)
    if i % 2:
        gaps.append((i, size, at_block))  # index, size-of-gap, location
        for _ in range(size):
            block_list.append('.')
    else:
        files.append((i // 2, size, at_block))  # id, num_blocks, location
        for _ in range(size):
            block_list.append(i // 2)
    at_block += size

# last = block_list.pop()
# print(block_list)
org_blocks = block_list.copy()
while True:
    last = block_list.pop()
    try:
        index = block_list.index('.')
    except ValueError:
        block_list.append(last)
        break
    block_list[index] = last

# print(block_list)

print(bchecksum(block_list))

block_list = org_blocks.copy()
# print(''.join(map(str, block_list)))
# print(map(str, block_list))
# done = set()
file_idx = len(files)
while file_idx >= 0:
    file_idx -= 1
    id, size, org_loc = files[file_idx]
    # done.add(id)
    # print(''.join(map(str, block_list)))
    # print(block_list)
    # print(f"file:{id}, size:{size}  ", end="")
    if '.' not in block_list[:org_loc]:
        print ("bkrea")
        break
    move_to = 0
    for i, gap in enumerate(gaps):
        # print(f"g:{gap}  ", end="")
        if gap[1] >= size:
            move_to = gap[2]
            gaps[i] = (gap[0], gap[1] - size, gap[2] + size)
            # print(f" moving to {move_to}")
            for jj in range(size):
                block_list[jj + move_to] = id
            for jj in range(org_loc, org_loc + size):
                block_list[jj] = '.'
            break
    # if left > 0:
    #     gaps[i]=left
    # else:
    #     gaps.pop(i)i
# if move_to == -1:
#     print(f"no gap for {id} size:({size})")
# else:

# print(''.join(map(str, block_list)))
# print(block_list)

print(bchecksum(block_list))

# sorted_files = sorted(files, key=cmp_to_key(mysorter), reverse=False)
# sorted_gaps = sorted(gaps)
# print(files, gaps)
# print(sorted_files)
# for i, gap in enumerate(sorted_gaps):
#     g_index = sorted
