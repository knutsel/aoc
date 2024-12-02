from aocd import get_data

data = get_data(year=2024, day=2).splitlines()


def safe(levels):
    for i in range(1, len(levels)):
        if not 1 <= abs(levels[i] - levels[i - 1]) <= 3:
            return 0

    if levels == sorted(levels) or levels == sorted(levels)[::-1]:
        return 1

    return 0


levels = []
for line in data:
    llevels = []
    for val in line.split():
        llevels.append(int(val))
    levels.append(llevels)

p1 = 0
p2 = 0
for lline in levels:
    p1_r = safe(lline)
    p1 += p1_r
    if p1_r == 0:
        for i in range(len(lline)):
            p2_r = safe(lline[:i] + lline[i + 1:])
            p2 += p2_r
            if p2_r == 1:
                break

print(p1)
print(p1 + p2)
