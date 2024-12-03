import re

from aocd import get_data

instr = get_data(year=2024, day=3)

def sum(instr):
    mults = re.findall(r'mul\(\d+,\d+\)', instr)
    sum = 0
    for mult in mults:
        nums = re.findall(r'\d+', mult)
        sum += int(nums[0]) * int(nums[1])
    return sum

print(sum(instr))
# groups = re.search(r'(.*)don\'t\(\)(.*)do\(\)(.*)', instr)
# print(sum(groups.group(1)), sum(groups.group(2))), sum(groups.group(3))
#
# dodontstring = re.sub(r'don\'t\(\).*do\(\)',  '', instr)
# print(f"{sum(groups.group(1))+sum(dodontstring)+sum(groups.group(3))}")

