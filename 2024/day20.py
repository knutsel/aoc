import heapq

from utils import get_input, print_grid


def is_inbounds(y, x):
    if x < 0 or y < 0 or x >= len(grid[0]) or y >= len(grid):
        return False
    return True

def neighbors(y, x):
    return [(y-1, x), (y+1, x), (y, x-1), (y, x+1)]

def make_g(grid):
    graph = {}
    start = end = ""
    innerwalls = set()
    steps = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    for y, _ in enumerate(grid):
        for x, _ in enumerate(grid[y]):
            if grid[y][x] == 'S':
                start = (y, x)
            if grid[y][x] == 'E':
                end = (y, x)
            nnodes = {}
            for step in steps:
                neighbor = (y + step[0], x + step[1])
                if is_inbounds(neighbor[0], neighbor[1]):
                    if grid[neighbor[0]][neighbor[1]] != '#':
                        nnodes[neighbor] = 1
                    elif neighbor[0] > 0 and neighbor[1] > 0:
                        innerwalls.add(neighbor)
            graph[(y, x)] = nnodes

    return graph, start, end, innerwalls


def dijkstra(graph, node):
    distances = {node: float('inf') for node in graph}
    distances[node] = 0
    came_from = {node: None for node in graph}
    queue = [(0, node)]

    while queue:
        current_distance, current_node = heapq.heappop(queue)
        for next_node, weight in graph[current_node].items():
            distance_temp = current_distance + weight
            if distance_temp < distances[next_node]:
                distances[next_node] = distance_temp
                came_from[next_node] = current_node
                heapq.heappush(queue, (distance_temp, next_node))
    return distances, came_from


data = get_input(for_example=False, day=20)

grid = []
for y, line in enumerate(data):  # use y, x in everything --> [line-no][char-on-line]
    grid.append(list(line))

graph, start, end, innerwalls = make_g(grid)
distances, came_from = dijkstra(graph, start)
current_node = end
numsteps = numturns = 0
prevdir = dir = "^"
while came_from[current_node] != start:
    numsteps += 1
    if prevdir != dir:
        numturns += 1
        prevdir = dir
    if current_node[0] > came_from[current_node][0]:
        dir = "v"
    elif current_node[0] < came_from[current_node][0]:
        dir = "^"
    elif current_node[1] < came_from[current_node][1]:
        dir = "<"
    elif current_node[1] > came_from[current_node][1]:
        dir = ">"
    current_node = came_from[current_node]
    grid[current_node[0]][current_node[1]] = dir

# print(numsteps)
print_grid(grid)

cheats = {}
for w in innerwalls:
    print(f"wall: {w}")
    n = neighbors(*w)
    print(n[0])
    cheat_value = -1
    if n[0] in distances and n[1] in distances:
        if distances[n[0]].is_integer() and  distances[n[1]].is_integer():
            cheat_value = abs(distances[n[0]] - distances[n[1]])-2
            print(f"w:{w} 01, cheat value: {cheat_value}")
    if n[2] in distances and n[3] in distances:
        if distances[n[2]].is_integer() and  distances[n[3]].is_integer():
            cheat_value = abs(distances[n[2]] - distances[n[3]])-2
            print(f"w:{w} 23, cheat value: {cheat_value}")

    if cheat_value != -1:
        if cheat_value not in cheats:
            cheats[cheat_value] = 1
        else:
            cheats[cheat_value] += 1

p1=0
for k, v in sorted(cheats.items(), key=lambda item: item[1]):
    print(f"{k}: {v}")
    if k >= 100:
        p1+= cheats[k]

print(p1)
    # left = (w[0] - 1, w[1])
    # right = (w[0] + 2, w[1])
    # if left in distances and right in distances:
    #     if distances[left].is_integer() and distances[right].is_integer():
    #         cheat_value = abs(distances[left] - distances[right])
    #         print(f"w: {w}, cheat value: {cheat_value}")
    # below = (w[0], w[1] - 1)
    # above = (w[0], w[1] + 2)
    # if below in distances and above in distances:
    #     if distances[below].is_integer() and distances[above].is_integer():
    #         cheat_value = abs(distances[below] - distances[above])
    #         print(f"w: {w}, cheat value: {cheat_value}")

# print(f"d:{distances[end]}")
