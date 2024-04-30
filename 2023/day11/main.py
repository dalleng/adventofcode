def main():
    galaxy = []
    with open("input.txt") as f:
        galaxy = f.read().splitlines()

    rows_to_expand = []
    for i, row in enumerate(galaxy):
        if '#' not in set(row):
            rows_to_expand.append(i)

    cols_to_expand = []
    for j in range(len(galaxy[0])):
        if '#' not in set([galaxy[i][j] for i in range(len(galaxy))]):
            cols_to_expand.append(j)

    print(f"{rows_to_expand=}")
    print(f"{cols_to_expand=}")

    galaxies = []
    for i in range(len(galaxy)):
        for j in range(len(galaxy[0])):
            if galaxy[i][j] == '#':
                galaxies.append((i, j))

    sum = 0
    for i in range(len(galaxies)):
        for j in range(i+1, len(galaxies)):
            from_row, from_col = galaxies[i]
            to_row, to_col = galaxies[j]
            for row in range(min(from_row, to_row), max(from_row, to_row)):
                if row in rows_to_expand:
                    sum += 2
                else:
                    sum += 1
            for col in range(min(from_col, to_col), max(from_col, to_col)):
                if col in cols_to_expand:
                    sum += 2
                else:
                    sum += 1
    print(f"{sum=}")

    sum = 0
    for i in range(len(galaxies)):
        for j in range(i+1, len(galaxies)):
            from_row, from_col = galaxies[i]
            to_row, to_col = galaxies[j]
            for row in range(min(from_row, to_row), max(from_row, to_row)):
                if row in rows_to_expand:
                    sum += 1_000_000
                else:
                    sum += 1
            for col in range(min(from_col, to_col), max(from_col, to_col)):
                if col in cols_to_expand:
                    sum += 1_000_000
                else:
                    sum += 1

    print(f"{sum=}")

if __name__ == '__main__':
    main()