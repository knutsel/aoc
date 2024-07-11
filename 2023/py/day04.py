import re

import utils


def intersection(lst1, lst2):
    return list(set(lst1) & set(lst2))


lines = utils.read_file('input04.txt')
score = 0
copies = [1] * len(lines)  # makes an array of len(lines) long with 1 as initial value
for i, line in enumerate(lines):
    all_ints = re.findall('\d+', line)
    inter = intersection(all_ints[1:11], all_ints[11:])
    this_score = 0
    if len(inter) > 0:
        this_score = 2 ** (len(inter) - 1)
    score += this_score
    for j in range(len(inter)):
        copies[i + j + 1] += copies[i]

print(score)
print(sum(copies))
