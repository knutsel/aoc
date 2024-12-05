from functools import cmp_to_key

from aocd import get_data

ordering = {}
rev = {}


def mysorter(item1, item2):
    if item1 == item2:
        return 0
    if item1 in ordering:
        for val in ordering[item1]:
            if item2 == val:
                return 1
    return -1
    # if item1 > item2:
    #     return 1
    # else:
    #     return -1


lines = get_data(year=2024, day=5)

# with open("example.txt") as file:
#     lines = file.read()

for oline in lines.split("\n\n")[0].splitlines():
    v1, v2 = map(int, oline.split("|"))
    if not v1 in ordering:
        ordering[v1] = []
    ordering[v1].append(v2)
    if not v2 in rev:
        rev[v2] = []
    rev[v2].append(v1)
    # print(v1, v2, ordering[v1])

updates = []
for i, uline in enumerate(lines.split("\n\n")[1].splitlines()):
    updates.append(list(map(int, uline.split(","))))

p1 = p2 = 0
wrong_ordered_updates = []
for update in updates:
    right_order = True
    for i, page in enumerate(update[:-1]):
        if page in ordering:
            if update[i + 1] not in ordering[page]:
                right_order = False

    # print(update)
    for i, page in enumerate(update[1:]):
        # print(i, page)
        if page in rev:
            if update[i + 1 - 1] not in rev[page]:
                # print(f" {update[i+1 - 1]} not in {rev[page]} ")
                right_order = False

    if right_order:
        # print(update, (len(update) - 1) / 2)
        p1 += update[int(round((len(update) - 1) / 2))]
    else:
        wrong_ordered_updates.append(update)

for update in wrong_ordered_updates:
    u1 = sorted(update, key=cmp_to_key(mysorter), reverse=False)
    # u1 = sorted(update,key=mysorter, reverse:True)
    print(f"update {update} -> {u1}")
    p2 += u1[int(round((len(update) - 1) / 2))]

print(p1, p2)
