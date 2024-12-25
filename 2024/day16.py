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
    dir = "<"
    distances = {node: float('inf') for node in graph}
    distances[node] = 0
    came_from = {node: None for node in graph}
    queue = [(0, node)]

    while queue:
        current_distance, current_node = heapq.heappop(queue)
        for next_node, weight in graph[current_node].items():
            # if dir == "<>" and next_node[1] != current_node[1]:
            #     weight += 1000
            #     dir = "v^"
            # elif dir == "v^" and next_node[0] != current_node[0]:
            #     weight += 1000
            #     dir = "<>"
            distance_temp = current_distance + weight
            if distance_temp < distances[next_node]:
                distances[next_node] = distance_temp
                came_from[next_node] = current_node
                heapq.heappush(queue, (distance_temp, next_node))
    return distances, came_from


data = get_input(for_example=True, day=16)

grid = []
for y, line in enumerate(data):  # use y, x in everything --> [line-no][char-on-line]
    grid.append(list(line))

graph, start, end = make_g(grid)
distances, came_from = dijkstra(graph, start)
print_grid(grid)
print(distances[end], came_from[end])
current_node = end
numsteps = 0
while came_from[current_node] != start:
    print(f"{current_node} ->{came_from[current_node]}")
    numsteps += 1
    current_node = came_from[current_node]
    grid[current_node[0]][current_node[1]] = 'X'

print(numsteps)
print_grid(grid)
