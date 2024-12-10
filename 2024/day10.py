import heapq

from utils import get_input, print_grid


def dijkstra(graph, node):
    distances = {node: float('inf') for node in graph}
    distances[node] = 0
    came_from = {node: None for node in graph}
    queue = [(0, node)]

    while queue:
        current_distance, current_node = heapq.heappop(queue)
        for next_node, weight in graph[current_node].items():
            # ## mod to dijkstra's algorithm for this puzzle
            if weight != 1:
                continue
            # ##
            distance_temp = current_distance + weight
            if distance_temp < distances[next_node]:
                distances[next_node] = distance_temp
                came_from[next_node] = current_node
                heapq.heappush(queue, (distance_temp, next_node))
    return distances, came_from


def is_inbounds(y, x):
    if x < 0 or y < 0 or x >= len(grid[0]) or y >= len(grid):
        return False
    return True


def make_g(grid):
    graph = {}
    trail_starts = []
    steps = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    for y, _ in enumerate(grid):
        for x, _ in enumerate(grid[y]):
            if grid[y][x] == 0:
                trail_starts.append((y, x))
            nnodes = {}
            for step in steps:
                neighbor = (y + step[0], x + step[1])
                if is_inbounds(neighbor[0], neighbor[1]):
                    nnodes[neighbor] = grid[neighbor[0]][neighbor[1]] - grid[y][x]
            graph[(y, x)] = nnodes

    return graph, trail_starts


data = get_input(for_example=False, day=10)

grid = []
for y, line in enumerate(data):  # use y, x in everything --> [line-no][char-on-line]
    grid.append(list(map(int, list(line))))

graph, trail_starts = make_g(grid)


p1 = 0
for s in trail_starts:
    distances, came_from = dijkstra(graph, s)
    for k in distances.keys():
        d = distances[k]
        if d == 9:
            p1 += 1

print(f"p1: {p1}")
