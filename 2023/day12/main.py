import re
from typing import Generator


def main():
    grid, counts = read_input("input.txt")
    # print(f"{grid=}")
    # print(f"{counts=}")

    total_count = 0
    for i in range(len(grid)):
        positions = get_unknown_positions(grid[i])
        count = 0
        # print(f"{grid[i]=}")
        # print(f"{counts[i]=}")
        for replacement in get_possible_replacements(grid[i], positions):
            # print(f"{replacement=}")
            if counts[i] == list(map(len, re.findall(r'#+', replacement))):
                # print("found suitable replacement")
                count += 1
        total_count += count

    print(f"{total_count=}")


def read_input(filename):
    with open(filename) as f:
        grid = []
        counts = []
        for line in f:
            row, g = line.split(" ")
            grid.append(row)
            counts.append([int(i) for i in g.split(",")])
    return grid, counts


def get_unknown_positions(row: str):
    positions = []
    for i, c in enumerate(row):
        if c == '?':
            positions.append(i)
    return positions


def get_possible_replacements(row: str, positions: list[int]) -> Generator[str, None, None]:
    to_process = [(row, positions)]
    while to_process:
        current_row, current_positions = to_process.pop()
        if not current_positions:
            yield current_row
        else:
            for c in (".", "#"):
                pos = current_positions[0]
                new_row = current_row[:pos] + c + current_row[pos+1:]
                to_process.append((new_row, current_positions[1:]))


if __name__ == "__main__":
    main()
