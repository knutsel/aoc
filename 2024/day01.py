from aocd import get_data

data = get_data(year=2024, day=1).splitlines()
ldata = []
rdata = []
for line in data:
    ldata.append(int(line.split()[0]))
    rdata.append(int(line.split()[1]))

ldata.sort()
rdata.sort()

dist = 0
sim = 0
for i, _ in enumerate(ldata):
    dist += abs(rdata[i] - ldata[i])
    count = rdata.count(ldata[i])
    sim += count * ldata[i]
    i += 1

print(f"part1: {dist} part2:{sim}")
