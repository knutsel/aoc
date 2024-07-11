import re

import utils

lines: list[str] = utils.read_file('example.txt')
seeds: list[int] = list(map(int, re.findall('\d+', lines[0])))
lines.pop(0)
current_key: str = "NOTRIGHT"
# rules: list[list[(int,int, int)]] = {}
# mapping: dict[str, str] = {}
rules: [[(int, int, int)]] = [[]]
conversion: [(int,int, int)] = []
# j = 0
for i, line in enumerate(lines):
    if not line:
        continue
    if m := re.match('(.*) map:', line):
        # print(m.group(1))
        # current_key = m.group(1)
        # ft = m.group(1).split('-to-')
        # mapping[ft[0]] = ft[1]
        if len(conversion)  != 0:
            rules.append(conversion)
        conversion = []
        continue
    else:
        # if current_key in rules:
        t = list(map(int, line.split(' ')))
        print(t)
        conversion.append(t)
        # j += 1
    # else:
    #     rules[i] = [map(int, line.split(' '))]

print(rules)


def follow(level, value):
    while level < 8:
        mapping_rules = rules[level]
        new_value = value
        for rule in mapping_rules:
            if rule[1] <= value < rule[2] + rule[1]:
                new_value = rule[0] + (value - rule[1])
        # print(f"{what}-to-{mapping[what]} ->mapping_rules {mapping_rules}, value {value}, new value {new_value}")
        value = follow(level + 1, new_value)
    return value


locs: [int] = []
for s in seeds:
    location = follow(0, s)
    # print(location)
    locs.append(location)
print(min(locs))
#
# range_to_check: [(int, int)] = []
# for i in range(0, len(seeds), 2):
#     start = seeds[i]
#     end = seeds[i + 1]
#     range_to_check.append((start, end))
#
# while range_to_check:
#     start, end = range_to_check.pop()
#     s1 = follow('seed', start)
#     e1 = follow('seed', end)
#     locs.append(s1)
#     locs.append(e1)
#     print(start, end, s1, e1)
#     if start != s1 or end != e1:
#         range_to_check.append((start, int(end/2)))
#         range_to_check.append((int(end/2), end))
#     # print(f"i:{i}")
#     # for j in range(seeds[i], seeds[i]+seeds[i+1]):
#     #     # print(f"j:{j} from {seeds[i]} to {seeds[1]+seeds[i]+1}")
#     #     # locs.append(follow('seed', j))
#     #
# print(min(locs))
