from utils import get_input, print_grid


def lstripstr(num):
    for i in range(len(num)):

        if num[i] != '0':
            res = num[i::];
            return res

    return "0"

def blink(stones):
    new_stones = []
    for i, s in enumerate(stones):
        if s == '0':
            new_stones.append('1')
        elif len(s)%2 == 0:
            # # new_stones.append('0')
            # s2 = s1 = ""
            # for j, c in enumerate(s):
            #     if j%2 == 0:
            #         s1+=c
            #     else:
            #         s2+=c
            new_stones.append(lstripstr(s[:len(s) // 2]))
            new_stones.append(lstripstr(s[len(s) // 2:]))
        else:
            new_stones.append(str(int(s)*2024))
    return new_stones

data = get_input(for_example=False, day=11)[0] # one line

stones = data.split(' ')
for b in range(25):
    # print(f"b:{b}, data:\n{data}")
    stones = blink(stones)

print(f"data: {len(stones)}")
