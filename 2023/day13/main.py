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
    transposed = ["".join([row[n] for row in pattern]) for n in range(len(pattern[0]))]
    return find_row_reflection(transposed)


def find_row_reflection(pattern: Pattern) -> int | None:
    for i in range(len(pattern) - 1):
        if pattern[i] == pattern[i + 1]:
            is_reflection = True
            reflection_size = min(i, (len(pattern) - 1) - (i + 1))

            for d in range(1, reflection_size + 1):
                if pattern[i - d] != pattern[i + 1 + d]:
                    is_reflection = False
                    break

            if is_reflection:
                return i + 1  # problem uses 1-based indexing


def equal_or_almost_equal(s1: str, s2: str, allowed_diff=1) -> tuple[bool, list[int]]:
    if len(s1) != len(s2):
        return False, []

    count_diff = 0
    indices = []

    for i in range(len(s1)):
        if s1[i] != s2[i]:
            count_diff += 1
            indices.append(i)

    return (True, indices) if count_diff <= allowed_diff else (False, indices)


def find_col_reflection_with_smudge(pattern: Pattern) -> int | None:
    transposed = ["".join([row[n] for row in pattern]) for n in range(len(pattern[0]))]
    return find_row_reflection_with_smudge(transposed)


def find_row_reflection_with_smudge(pattern: Pattern) -> int | None:
    for i in range(len(pattern) - 1):
        equal, diff_indices = equal_or_almost_equal(pattern[i], pattern[i + 1])

        if equal:
            smudge_found = bool(diff_indices)
            is_reflection = True
            reflection_size = min(i, (len(pattern) - 1) - (i + 1))

            for d in range(1, reflection_size + 1):
                allowed_diff = 0 if smudge_found else 1
                equal, diff_indices = equal_or_almost_equal(pattern[i-d], pattern[i+1+d], allowed_diff=allowed_diff)

                if not smudge_found:
                    smudge_found = bool(diff_indices)

                if not equal:
                    is_reflection = False
                    break

            if is_reflection and smudge_found:
                return i + 1  # problem uses 1-based indexing


def main():
    patterns = read_input("input.txt")
    s1 = 0
    s2 = 0

    for p in patterns:
        if col := find_col_reflection(p):
            s1 += col

        if col := find_col_reflection_with_smudge(p):
            s2 += col

        if row := find_row_reflection(p):
            s1 += row * 100

        if row := find_row_reflection_with_smudge(p):
            s2 += row * 100

    print(f"{s1=}")
    print(f"{s2=}")


if __name__ == "__main__":
    import sys

    print(f"{sys.argv=}")

    if len(sys.argv) > 1 and sys.argv[1] == "--test":
        print("Running tests")
        patterns = read_input("input_example.txt")
        assert find_col_reflection(patterns[0]) == 5
        assert find_col_reflection(patterns[1]) is None
        assert find_col_reflection(patterns[2]) == 5
        assert find_row_reflection(patterns[0]) is None
        assert find_row_reflection(patterns[1]) == 4
        assert find_row_reflection(patterns[2]) is None

        assert equal_or_almost_equal("a", "b", allowed_diff=0) == (False, [0])
        assert equal_or_almost_equal("", "", allowed_diff=0) == (True, [])
        assert equal_or_almost_equal("", "") == (True, [])
        assert equal_or_almost_equal("a", "b") == (True, [0])
        assert equal_or_almost_equal("a", "b", allowed_diff=2) == (True, [0])
        assert equal_or_almost_equal("aa", "bb", allowed_diff=2) == (True, [0, 1])
        assert equal_or_almost_equal("ab", "ac") == (True, [1])
        assert equal_or_almost_equal("ab", "bb") == (True, [0])

        assert find_row_reflection_with_smudge(patterns[0]) == 3
        assert find_col_reflection_with_smudge(patterns[0]) is None

        assert find_row_reflection_with_smudge(patterns[1]) == 1
        assert find_col_reflection_with_smudge(patterns[1]) is None
        assert find_row_reflection_with_smudge(patterns[2]) == 3
        assert find_col_reflection_with_smudge(patterns[2]) is None
    else:
        main()
