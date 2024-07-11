import utils

lines = utils.read_file('input01.txt')

# Part 1
total = 0
for line in lines:
    right = -1
    left = -1
    for c in line:
        if c.isnumeric():
            if left == -1:
                left = int(c)
            right = int(c)
    val = left * 10 + right
    total += val
print(f"part1 {total}")

# Part 2
numbers = dict(one="o1e", two="t2o", three="t3e", four="f4r", five="f5e", six="s6x", seven="s7n", eight="e8t",
               nine="n9e")

total = 0
for line in lines:
    right = -1
    left = -1
    for num in numbers:
        line = line.replace(num, numbers[num])
    for c in line:
        if c.isnumeric():
            if left == -1:
                left = int(c)
            right = int(c)
    val = left * 10 + right
    total += val
print(f"part2 {total}")
