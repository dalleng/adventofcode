import re
from collections import Counter


def main():
    s = 0
    s2 = 0
    counter = Counter()

    with open('input.txt') as file:
        cards = file.read().split('\n')

    for i, line in enumerate(cards, start=1):
        counter[i] += 1
        print(f"card {i} has {counter[i]} copies")

        winning_numbers, own_numbers = extract_numbers(line)

        intersection = set(winning_numbers).intersection(set(own_numbers))
        matching_cards = len(intersection)

        for j in range(1, matching_cards+1):
            counter[i+j] += 1 * counter[i]

        s += 2 ** (matching_cards-1) if intersection else 0

    s2 = sum(counter.values())
    print(f"{s=}")
    print(f"{s2=}")


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
