import re
from typing import NamedTuple
from collections import defaultdict


class PartNumber(NamedTuple):
    row: int
    start: int
    end: int
    number: int


def main():
    with open("input.txt") as f:
        schematic = [line.strip() for line in f]

    s1 = 0
    adjacent_asterisks = defaultdict(list)

    for part in find_parts(schematic):
        if is_valid(schematic, part, adjacent_asterisks):
            s1 += part.number

    print(f"{s1=}")

    s2 = 0
    for _, value in adjacent_asterisks.items():
        if len(value) == 2:
            s2 += value[0] * value[1]

    print(f"{s2=}")


def find_parts(schematic):
    for row, line in enumerate(schematic):
        for m in re.finditer(r"\d+", line):
            start, end = m.span()
            yield PartNumber(
                number=int(m.group()),
                row=row,
                start=start,
                end=end-1
            )


def is_valid(schematic, part, adjacent_asterisks):
    is_valid = False

    # check one position below start
    row = part.row
    col = part.start - 1
    if part.start > 0 and is_symbol(schematic[row][col]):
        if schematic[row][col] == '*':
            adjacent_asterisks[(row, col)].append(part.number)
        is_valid = True

    # check the upper diagonal of the start
    row = part.row - 1
    col = part.start - 1
    if (
        part.start > 0
        and part.row > 0
        and is_symbol(schematic[row][col])
    ):
        if schematic[row][col] == '*':
            adjacent_asterisks[(row, col)].append(part.number)
        is_valid = True

    # check the lower diagonal of the start
    row = part.row + 1
    col = part.start - 1
    if (
        part.start > 0
        and part.row < len(schematic) - 1
        and is_symbol(schematic[row][col])
    ):
        if schematic[row][col] == '*':
            adjacent_asterisks[(row, col)].append(part.number)
        is_valid = True

    row = part.row
    col = part.end + 1
    # check one postion after the end
    if part.end < len(schematic[0]) - 1 and is_symbol(
        schematic[row][col]
    ):
        if schematic[row][col] == '*':
            adjacent_asterisks[(row, col)].append(part.number)
        is_valid = True

    # check upper diagonal of the end
    row = part.row - 1
    col = part.end + 1
    if (
        part.end < len(schematic[0]) - 1
        and part.row > 0
        and is_symbol(schematic[row][col])
    ):
        if schematic[row][col] == '*':
            adjacent_asterisks[(row, col)].append(part.number)
        is_valid = True

    # check the lower diagonal of the end
    row = part.row + 1
    col = part.end + 1
    if (
        part.end < len(schematic[0]) - 1
        and part.row < len(schematic) - 1
        and is_symbol(schematic[row][col])
    ):
        if schematic[row][col] == '*':
            adjacent_asterisks[(row, col)].append(part.number)
        is_valid = True

    for i in range(part.start, part.end+1):
        # check above
        row = part.row - 1
        if part.row > 0 and is_symbol(schematic[row][i]):
            is_valid = True
            if schematic[row][i] == '*':
                adjacent_asterisks[(row, i)].append(part.number)

        # check below
        row = part.row + 1
        if part.row < len(schematic) - 1 and is_symbol(schematic[row][i]):
            is_valid = True
            if schematic[row][i] == '*':
                adjacent_asterisks[(row, i)].append(part.number)

    return is_valid, adjacent_asterisks


def is_symbol(c):
    return not c.isdigit() and c != "."


if __name__ == "__main__":
    # assert is_symbol("@") is True
    main()
