import re


def get_path_length(
        map: dict[str, tuple[str, str]],
        movements: list[str],
        origin: str = 'AAA',
        destination: str = 'ZZZ'
):
    current_position = origin
    steps = 0
    while current_position != destination:
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

    # ending_positions = {key for key in map.keys() if key[2] == "Z"}
    # print(f"{ending_positions=}")
    # print(f"{len(set(map.keys()))=}")
    # print(f"{len(list(map.keys()))=}")

    # final_positions = build_final_positions_map(map, movements)
    # print(f"s2={get_path_length2(final_positions, movements)}")


if __name__ == "__main__":
    main()
