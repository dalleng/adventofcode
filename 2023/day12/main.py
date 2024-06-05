from functools import cache


def main():
    grid, counts = read_input("input.txt")
    # print(f"{grid=}")
    # print(f"{counts=}")

    s1 = 0
    s2 = 0

    for i in range(len(grid)):
        s1 += count(grid[i], counts[i])
        s2 += count('?'.join(grid[i] for _ in range(5)), counts[i] * 5)

    print(f"{s1=}")
    print(f"{s2=}")


@cache
def count(springs: str, counts: tuple[int, ...]) -> int:
    if springs == "":
        return 1 if counts == () else 0

    if counts == ():
        return 1 if "#" not in springs else 0

    results = 0

    if springs[0] in ".?":
        results += count(springs[1:], counts)

    if springs[0] in "?#":
        if (
            counts[0] <= len(springs)
            and "." not in springs[: counts[0]]
            and (len(springs) == counts[0] or springs[counts[0]] in ".?")
        ):
            results += count(springs[counts[0]+1:], counts[1:])

    return results


def read_input(filename):
    with open(filename) as f:
        grid = []
        counts = []
        for line in f:
            row, c = line.split(" ")
            grid.append(row)
            counts.append(tuple(map(int, c.split(","))))

    return grid, counts


if __name__ == "__main__":
    main()
