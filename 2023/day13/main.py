Pattern = list[str]


def read_input(filepath: str) -> list[Pattern]:
    patterns = [[]]
    with open(filepath) as f:
        for line in f:
            if stripped := line.strip():
                patterns[-1].append(stripped)
            else:
                patterns.append([])
    return patterns


def find_col_reflection(pattern: Pattern) -> int | None:
    transposed = [''.join([row[n] for row in pattern]) for n in range(len(pattern[0]))]
    return find_row_reflection(transposed)


def find_row_reflection(pattern: Pattern) -> int | None:
    for i in range(len(pattern) - 1):
        if pattern[i] == pattern[i + 1]:
            is_reflection = True
            reflection_size = min(i, (len(pattern)-1) - (i+1))

            for d in range(1, reflection_size+1):
                if pattern[i-d] != pattern[i+1+d]:
                    is_reflection = False

            if is_reflection:
                return i + 1  # problem uses 1-based indexing


def main():
    patterns = read_input("input.txt")
    s1 = 0
    for p in patterns:
        if col := find_col_reflection(p):
            s1 += col

        if row := find_row_reflection(p):
            s1 += row * 100

    print(f"{s1=}")


if __name__ == "__main__":
    main()
    # patterns = read_input("input_example.txt")
    # assert find_col_reflection(patterns[0]) == 5
    # assert find_col_reflection(patterns[1]) is None
    # assert find_col_reflection(patterns[2]) == 5
    # assert find_row_reflection(patterns[0]) is None
    # assert find_row_reflection(patterns[1]) == 4
    # assert find_row_reflection(patterns[2]) is None
