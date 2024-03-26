import math
import re
from typing import Callable


def get_path_length(
        map: dict[str, tuple[str, str]],
        movements: list[str],
        origin: str = 'AAA',
        is_destination_fn: Callable[[str], bool] = lambda x: x == "ZZZ",
):
    current_position = origin
    steps = 0
    while not is_destination_fn(current_position):
        for mov in movements:
            direction = 0 if mov == "L" else 1
            current_position = map[current_position][direction]
            steps += 1
    return steps


def main():
    with open("input.txt") as f:
        map = {}
        movements = next(f).strip()

        for line in f:
            if not line:
                continue

            m = re.match(r"([A-Z1-9]{3}) = \(([A-Z1-9]{3}), ([A-Z1-9]{3})\)", line)

            if m:
                origin, left, right = m.groups()
                map[origin] = (left, right)

    s1 = get_path_length(map, movements)
    assert s1 == 15989
    print(f"{s1=}")

    starting_positions = [pos for pos in map.keys() if pos[2] == 'A']
    path_lengths = []

    for pos in starting_positions:
        pl = get_path_length(
            map, movements, origin=pos, is_destination_fn=lambda x: x[2] == 'Z'
        )
        path_lengths.append(pl)

    # Find the lowest common multiple of all the path lengths
    # of starting positions to positions ending with 'Z'
    print(f"{path_lengths=}")
    s2 = math.lcm(*path_lengths)
    assert s2 == 13830919117339
    print(f"{s2=}")


if __name__ == "__main__":
    main()
