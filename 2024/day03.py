import re

from aocd import get_data

instr = get_data(year=2024, day=3)


def sum(instr):
    mults = re.findall(r'mul\(\d+,\d+\)', instr, re.MULTILINE)
    sum = 0
    for mult in mults:
        nums = re.findall(r'\d+', mult)
        sum += int(nums[0]) * int(nums[1])
    return sum


print(sum(instr))

s = re.subn(r'don\'t\(\).*?do\(\)', '_____', instr, 0, re.MULTILINE)
# print(s[0])
print(s[1], sum(s[0]))

# groups = re.findall(r'(do\(\).*?)(don\'t\(\).*?)', "do()"+instr+"don't()")
#
# p2=0
# for g in groups:
#     print(g[0])
#     p2+=sum(g[0])
#
# print(p2)
# groups = re.search(r'(.*?)don\'t\(\)(.*)do\(\)(.*?)', instr)
# # print(sum(groups.group(1)), sum(groups.group(2))), sum(groups.group(3))
# #
# dodontstring = re.sub(r'don\'t\(\).*do\(\)',  '', instr)
# print(groups.groups(1)[0])
# print(groups.groups(1)[1])
# print(f"{sum(groups.group(1)[0])+sum(groups.group(1)[1])+sum(groups.group(1)[2])}")
