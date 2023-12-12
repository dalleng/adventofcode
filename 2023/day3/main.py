import re
from typing import NamedTuple


class PartNumber(NamedTuple):
    row: int
    start: int
    end: int
    number: int


def main():
    with open("input.txt") as f:
        schematic = [line.strip() for line in f]
        # print(schematic)

    s = 0
    for part in find_parts(schematic):
        print(f"{part=}")
        if is_valid(schematic, part):
            print("is valid")
            s += part.number
        else:
            print("is not valid")

    print(f"{s=}")


def find_parts(schematic):
    for row, line in enumerate(schematic):
        for m in re.finditer(r"\d+", line):
            start, end = m.span()
            yield PartNumber(number=int(m.group()), row=row, start=start, end=end-1)


def is_valid(schematic, part):
    # check one position below start
    if part.start > 0 and is_symbol(schematic[part.row][part.start - 1]):
        return True

    # check the upper diagonal of the start
    if (
        part.start > 0
        and part.row > 0
        and is_symbol(schematic[part.row - 1][part.start - 1])
    ):
        return True

    # check the lower diagonal of the start
    if (
        part.start > 0
        and part.row < len(schematic) - 1
        and is_symbol(schematic[part.row + 1][part.start - 1])
    ):
        return True

    # check one postion after the end
    if part.end < len(schematic[0]) - 1 and is_symbol(
        schematic[part.row][part.end + 1]
    ):
        return True

    # check upper diagonal of the end
    if (
        part.end < len(schematic[0]) - 1
        and part.row > 0
        and is_symbol(schematic[part.row - 1][part.end + 1])
    ):
        return True

    # check the lower diagonal of the end
    if (
        part.end < len(schematic[0]) - 1
        and part.row < len(schematic) - 1
        and is_symbol(schematic[part.row + 1][part.end + 1])
    ):
        return True

    for i in range(part.start, part.end+1):
        # check above
        if part.row > 0 and is_symbol(schematic[part.row - 1][i]):
            return True

        # check below
        if part.row < len(schematic) - 1 and is_symbol(schematic[part.row + 1][i]):
            return True


def is_symbol(c):
    return not c.isdigit() and c != "."


if __name__ == "__main__":
    # assert is_symbol("@") is True
    main()
