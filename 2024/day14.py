import os
import re
import time

from utils import get_input, print_grid

in_lines = get_input(for_example=False, day=14)
wide = 101
tall = 103


# wide = 7
# tall = 11


class RobotClass:
    def __init__(self, px, py, vx, vy):
        self.px = px
        self.py = py
        self.vx = vx
        self.vy = vy
        self.start = (px, py)
        self.moves_done = 0
        self.back_at_start_after = 0

    def move(self):
        self.px = (self.px + self.vx) % wide
        self.py = (self.py + self.vy) % tall
        # self.px += self.vx
        # self.py += self.vy
        self.moves_done += 1
        if (self.px, self.py) == self.start:
            print("BACK")
            self.back_at_start_after = self.moves_done

    def quadrant(self):
        if self.px > wide // 2:
            if self.py > tall // 2:
                return 2
            elif self.py < tall // 2:
                return 1
        elif self.px < wide // 2:
            if self.py > tall // 2:
                return 3
            elif self.py < tall // 2:
                return 0


robots = []
for l in in_lines:
    try:
        groups = re.findall(r'^p=(\d+),(\d+) v=(-?\d*\.{0,1}\d+),(-?\d*\.{0,1}\d+)$', l)
        robots.append(RobotClass(int(groups[0][0]), int(groups[0][1]), int(groups[0][2]), int(groups[0][3])))
    except IndexError:
        print("HA")

p1 = p2 = 0

for i in range(10000):
    grid = [[0 for col in range(wide)] for row in range(tall)]
    for r in robots:
        r.move()
        grid[r.py][r.px] = 1
        # if i % 2 == 0:
        # print('\033[2J')
    os.system('clear')
    print_grid(grid)
    time.sleep(.01)
    print(i)
    # inp = input()

quadrants = [0, 0, 0, 0]
for r in robots:
    if r.quadrant() != None:
        quadrants[r.quadrant()] += 1

print(quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3], p2)
