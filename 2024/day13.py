import re

from utils import get_input

in_lines = get_input(for_example=False, day=13)


class MachineClass:
    A = (0, 0)
    B = (0, 0)
    P = (0, 0)

    def __init__(self, price_a, price_b):
        self.price_a = price_a
        self.price_b = price_b

    def __str__(self):
        return f"Button A: {self.A}, Button B: {self.B}, P: {self.P}, (prA: {self.price_a}, prB: {self.price_b})"

    def solve(self):
        # from the internet, math is hard, I'm doing this for fun.
        # A = (p_x * b_y - prize_y * b_x) / (a_x * b_y - a_y * b_x)
        # B = (a_x * p_y - a_y * p_x) / (a_x * b_y - a_y * b_x)
        SA = (self.P[0] * self.B[1] - self.P[1] * self.B[0]) / (self.A[0] * self.B[1] - self.A[1] * self.B[0])
        SB = (self.A[0] * self.P[1] - self.A[1] * self.P[0]) / (self.A[0] * self.B[1] - self.A[1] * self.B[0])
        if SB.is_integer() and SA.is_integer():
            return int(SA*self.price_a + SB*self.price_b)
        else:
            return 0

    def solve2(self):
        P2=(self.P[0]+10000000000000, self.P[1]+10000000000000)
        SA = (P2[0] * self.B[1] - P2[1] * self.B[0]) / (self.A[0] * self.B[1] - self.A[1] * self.B[0])
        SB = (self.A[0] * P2[1] - self.A[1] * P2[0]) / (self.A[0] * self.B[1] - self.A[1] * self.B[0])
        if SB.is_integer() and SA.is_integer():
            return int(SA * self.price_a + SB * self.price_b)
        else:
            return 0


machines = []  #
mach_no = 0
m = MachineClass(3, 1)
for l in in_lines:
    groups = re.findall(r'^(.*:).*X.(\d+), Y.(.+)$', l)
    if len(groups) == 0:
        machines.append(m)
        m = MachineClass(3, 1)
        continue
    match groups[0][0]:
        case ('Button A:'):
            m.A = (int(groups[0][1]), int(groups[0][2]))
        case ('Button B:'):
            m.B = (int(groups[0][1]), int(groups[0][2]))
        case ('Prize:'):
            m.P = (int(groups[0][1]), int(groups[0][2]))
machines.append(m)

p1 = p2= 0
for m in machines:
    p1+=m.solve()
    p2+=m.solve2()
    # print(m)

print(p1, p2)
