import heapq

from utils import get_input, print_grid


def is_inbounds(y, x):
    if x < 0 or y < 0 or x >= len(grid[0]) or y >= len(grid):
        return False
    return True


def make_g(grid):
    graph = {}
    steps = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    for y, _ in enumerate(grid):
        for x, _ in enumerate(grid[y]):
            nnodes = {}
            for step in steps:
                neighbor = (y + step[0], x + step[1])
                if is_inbounds(neighbor[0], neighbor[1]):
                    if grid[neighbor[0]][neighbor[1]] != '#':
                        nnodes[neighbor] = 1
            graph[(y, x)] = nnodes

    return graph, (len(grid)-1, len(grid[0])-1), (0,0)


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


example = False
data = get_input(for_example=example, day=18)
stop = 1024
numrows = numcols = 71
if example:
    numrows = numcols = 7
    stop = 12

grid = [["." for x in range(numcols)] for y in range(numrows)]

blocked = []
for i in range(len(data)):  # use y, x in everything --> [line-no][char-on-line]
    line = data[i]
    r, c = int(line.split(",")[0]), int(line.split(",")[1])
    blocked.append((r, c))
    grid[r][c] = '#'
    if i < 1024:
        continue
    graph, start, end = make_g(grid)
    distances, came_from = dijkstra(graph, start)
    if i == 1024:
        print(f" p1: {distances[end]}")
    if distances[end] > 9999999:
        print(f"{i} {data[i]} -> {distances[end]}")
        break

