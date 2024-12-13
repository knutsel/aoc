from utils import get_input


def lstripstr(num):
    for i in range(len(num)):

        if num[i] != '0':
            res = num[i::];
            return res

    return "0"


l1cache = {}

cache = {}


# def blink_recursive(s, blinks):
#     new_val = ''
#     if blinks == 0:
#         return 1
#     if (s, blinks) in cache:
#         return cache[(s, blinks)]
#     if s == '0':
#         new_val = blink_recursive('1', blinks - 1)
#     elif len(s) % 2 == 0:
#         new_val = blink_recursive(lstripstr(s[:len(s) // 2]), blinks - 1) + blink_recursive(lstripstr(s[len(s) // 2:]),
#                                                                                             blinks - 1)
#     else:
#         new_val = blink_recursive(str(int(s) * 2024), blinks - 1)
#     cache[(s, blinks)] = new_val
#     return new_val
memory = {}

def blink_recursive(stone, blinks):
    if blinks == 0:
        return 1
    elif (stone, blinks) in memory:
        return memory[(stone, blinks)]
    elif stone == 0:
        val = blink_recursive(1, blinks - 1)
    elif len(str_stone := str(stone)) % 2 == 0:
        mid = len(str_stone) // 2
        val = blink_recursive(int(str_stone[:mid]), blinks - 1) + blink_recursive(int(str_stone[mid:]), blinks - 1)
    else:
        val = blink_recursive(stone * 2024, blinks - 1)
    memory[(stone, blinks)] = val
    return val



def blink(stones):
    new_stones = []
    for i, s in enumerate(stones):
        if s in l1cache:
            # print("HIT")
            new_val = l1cache[s]
        elif s == '0':
            new_val = ['1']
        elif len(s) % 2 == 0:
            new_val = [lstripstr(s[:len(s) // 2]), lstripstr(s[len(s) // 2:])]
        else:
            new_val = [str(int(s) * 2024)]
        l1cache[s] = new_val
        new_stones.extend(new_val)

    return new_stones


data = get_input(for_example=False, day=11)[0]  # one line

# cache = {}

stones = data.split(' ')
# for b in range(25):
#     print(f"b:{b}, len:\n{len(stones)}")
#     # num_stones = len(stones) - 1
#     # while num_stones > 0:
#     #     # if ':'.join(stones[:num_stones]) in cache:
#     #     #     print(f"hit! {num_stones}")
#     #     num_stones -= 1
#     nstones = blink(stones)
#     # cache[':'.join(stones)] = nstones
#     stones = nstones

print(f"data: {len(stones)}")
snums = list(map(int,stones))
p2r = p1r = 0
for s in snums:
    p1r += blink_recursive(s, 25)
    p2r += blink_recursive(s, 75)
print(f"recursive: p1: {p1r}, p2: {p2r}")
# print(sum(blink_recursive(stone, 25) for stone in snums))