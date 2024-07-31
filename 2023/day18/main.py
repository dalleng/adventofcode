# 1. execute all movements and store edge coordinates
# 2. get the minimun coordinate (top left coordinate)
# 3. calculate the south east (right down) diagonal of the coordinate from the previous step
# 4. flood fill with the coordinate from the previous step

def main():
    with open("input.txt") as f:
        moves = [tuple(line.strip().split()) for line in f.read().splitlines()]

    edges = get_edge_coordinates(moves)
    start = (min(edges)[0] + 1, min(edges)[1] + 1)
    all = flood_fill(start, edges)

    print(f"s1={len(all)}")


def get_edge_coordinates(moves: list[tuple]) -> list[tuple[int, int]]:
    direction_to_pos_diff = {"U": (-1, 0), "D": (1, 0), "R": (0, 1), "L": (0, -1)}
    edges = [(0, 0)]

    for direction, steps, _ in moves:
        steps = int(steps)
        diff = direction_to_pos_diff[direction]
        for _ in range(steps):
            last = edges[-1]
            new_coord = (last[0] + diff[0], last[1] + diff[1])
            if new_coord != (0, 0):
                edges.append(new_coord)

    return edges


def get_neighbors(pos: tuple[int, int]) -> list[tuple[int, int]]:
    neighbors = []
    for rdiff in (-1, 0, 1):
        for cdiff in (-1, 0, 1):
            neighbor = (pos[0] + rdiff, pos[1] + cdiff)
            if neighbor != pos:
                neighbors.append(neighbor)
    return neighbors


def flood_fill(start: tuple[int, int], edges: list[tuple[int, int]]) -> list[tuple[int, int]]:
    """
    If you start from a coordinate inside and avoid expanding to positions
    on the edge you'll always remain inside the area.
    """

    # coordinates holds the edges and already visited coordinates
    coordinates = set(edges)
    frontier = [start]

    while frontier:
        current = frontier.pop()

        if current in coordinates:
            continue

        coordinates.add(current)

        for neighbor in get_neighbors(current):
            if neighbor not in coordinates:
                frontier.append(neighbor)

    return list(coordinates)


if __name__ == "__main__":
    main()
