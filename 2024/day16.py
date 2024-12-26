import heapq

from utils import get_input, print_grid


def is_inbounds(y, x):
    if x < 0 or y < 0 or x >= len(grid[0]) or y >= len(grid):
        return False
    return True


def make_g(grid):
    graph = {}
    start = end = ""
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
            graph[(y, x)] = nnodes

    return graph, start, end


def dijkstra(graph, node):
    distances = {node: float('inf') for node in graph}
    distances[node] = 0
    came_from = {node: None for node in graph}
    queue = [(0, node)]

    while queue:
        current_distance, current_node = heapq.heappop(queue)
        for next_node, weight in graph[current_node].items():
            prev = came_from[current_node]
            if prev is None:
                prev = (node[0], node[1], +1)  # start East facing
            if not (prev[0] == current_node[0] == next_node[0]) and not (prev[1] == current_node[1] == next_node[1]):
                weight += 1000
            distance_temp = current_distance + weight
            if distance_temp < distances[next_node]:
                distances[next_node] = distance_temp
                came_from[next_node] = current_node
                heapq.heappush(queue, (distance_temp, next_node))
    return distances, came_from


data = get_input(for_example=False, day=16)

grid = []
for y, line in enumerate(data):  # use y, x in everything --> [line-no][char-on-line]
    grid.append(list(line))

graph, start, end = make_g(grid)
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

print(numsteps)
print_grid(grid)

# the distance works for the examples, but is off-by-four for the real thing :-(
print(distances[end], numsteps, numturns, numsteps + 1000 * numturns)
