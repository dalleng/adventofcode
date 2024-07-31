import re


def main():
    # with open("example.txt") as f:
    with open("input.txt") as f:
        moves = [tuple(line.strip().split()) for line in f.read().splitlines()]

    # Solution Part 1
    # 1. execute all movements and store edge coordinates
    # 2. get the minimun coordinate (top left coordinate)
    # 3. calculate the south east (right down) diagonal of the coordinate from the previous step
    # 4. flood fill with the coordinate from the previous step

    edges = get_edge_coordinates(moves)
    start = (min(edges)[0] + 1, min(edges)[1] + 1)
    all = flood_fill(start, edges)
    print(f"s1={len(all)}")

    # Solution Part 2
    moves2 = []
    for _, __, hexvalue in moves:
        hexvalue = re.sub(r'[\(\)#]', '', hexvalue)
        direction = {'0': 'R', '1': 'D', '2': 'L', '3': 'U'}[hexvalue[-1]]
        steps = int(hexvalue[:-1], base=16)
        moves2.append((direction, steps, None))

    vertices = get_vertex_coordinates(moves2)

    area = 0
    # Compute the shoe lace formula using the vertices of the loop
    # https://en.wikipedia.org/wiki/Shoelace_formula
    # https://brilliant.org/wiki/area-of-a-polygon/
    for i in range(len(vertices)-1):
        x1, y1 = vertices[i]
        x2, y2 = vertices[i+1]
        area += (x1*y2 - x2*y1) / 2

    area = abs(area)
    perimeter = sum(steps for _, steps, __ in moves2)

    """
    Why do we add perimeter/2 + 1?

    - Adding perimeter/2 is needed because each point in the perimeter is actually
    a full 1 square meter block. If we consider perimeter coordinates to be in the
    middle of each of block we're missing half block of area.

    - The +1 is because of turns, since it is a loop there will be 4 more inward turns.
    For those 4, the shoelace area only accounts for a quarter of the area of the block,
    the perimeter / 2 adjustment adds half a block, but we're still missing a quarter.
    4 * 1/4 = 1.
    """

    area += perimeter / 2 + 1
    print(f"s2={area}")


def get_vertex_coordinates(moves: list[tuple]) -> list[tuple[int, int]]:
    direction_to_pos_diff = {"U": (-1, 0), "D": (1, 0), "R": (0, 1), "L": (0, -1)}
    vertices = [(0, 0)]

    for direction, steps, _ in moves:
        steps = int(steps)
        diff = direction_to_pos_diff[direction]
        last = vertices[-1]
        new_coord = (last[0] + steps*diff[0], last[1] + steps*diff[1])
        vertices.append(new_coord)

    return vertices


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
        for neighbor in get_neighbors(current):
            if neighbor not in coordinates:
                frontier.append(neighbor)
                coordinates.add(neighbor)

    return list(coordinates)


if __name__ == "__main__":
    main()
