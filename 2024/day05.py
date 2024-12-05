from aocd import get_data

lines = get_data(year=2024, day=5)

# with open("example.txt") as file:
#     lines = file.read()

ordering = {}
rev = {}
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

# print(rev)
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
            if update[i+1 - 1] not in rev[page]:
                # print(f" {update[i+1 - 1]} not in {rev[page]} ")
                right_order = False

    if right_order:
        # print(update, (len(update) - 1) / 2)
        p1 += update[int(round((len(update) - 1) / 2))]

print(p1, p2)
