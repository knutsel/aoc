import utils

lines = utils.read_file('input02.txt')
maxpart1 = {"red": 12, "green": 13, "blue": 14}
possible = 0
game_number = 1
power_sum = 0
for line in lines:
    max = {"red": 0, "green": 0, "blue": 0}
    for set in line.split(':')[1].split(';'):
        for cval in set.split(','):
            num, color = cval.lstrip().split(' ')
            if max[color] < int(num):
                max[color] = int(num)
    if max["red"] <= maxpart1["red"] and max["blue"] <= maxpart1["blue"] and max["green"] <= maxpart1["green"]:
        possible += game_number
    power_sum = power_sum + max["red"] * max["green"] * max["blue"]
    game_number = game_number + 1

print(f'part1: {possible} part2: {power_sum}')
