import re


def main():
    s = 0
    with open('input.txt') as file:
        for line in file:
            winning_numbers, own_numbers = extract_numbers(line)
            intersection = set(winning_numbers).intersection(set(own_numbers))
            s += 2 ** (len(intersection)-1) if intersection else 0
    print(f"{s=}")


def extract_numbers(line):
    line = line.strip()
    m = re.match(r'Card\s+\d+:([\d\s]*)\|([\d\s]*)', line)
    winning, own = m.groups()
    winning = [int(w) for w in winning.strip().split()]
    own = [int(o) for o in own.strip().split()]

    print(f"{winning=}")
    print(f"{own=}")

    return winning, own


if __name__ == "__main__":
    main()
