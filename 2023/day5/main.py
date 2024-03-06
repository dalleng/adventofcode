import re
import bisect


class RangeMap:
    def __init__(self, range_list):
        self.range_list = sorted(range_list)
        self.length = len(range_list)

    def __getitem__(self, key):
        i = bisect.bisect(self.range_list, key, key=lambda x: x[0])
        if i == 0:
            return key
        else:
            source, destination, size = self.range_list[i-1]
            if key >= source and key < source + size:
                return key - source + destination
        return key


def main():
    seed_numbers, range_maps = parse_input("input.txt")
    locations = []
    for seed in seed_numbers:
        current = seed
        for rm in range_maps:
            current = rm[current]
        locations.append(current)
    print(f"{locations=}")
    s1 = min(locations)
    print(f"{s1=}")


def parse_input(input_file):
    with open(input_file) as f:
        seeds = next(f)
        _, seed_numbers = seeds.split('seeds: ')
        seed_numbers = [int(s) for s in seed_numbers.split(' ')]

        range_maps = []
        range_list = []
        for line in f:
            if not f:
                continue

            m = re.match(r'([a-z]+)-to-([a-z]+)', line)
            if m:
                if range_list:
                    d = RangeMap(range_list)
                    range_maps.append(d)
                    range_list = []

            m = re.match(r'(\d+) (\d+) (\d+)', line)
            if m:
                destination, source, size = m.groups()
                range_list.append((int(source), int(destination), int(size)))

        d = RangeMap(range_list)
        range_maps.append(d)

    return seed_numbers, range_maps


if __name__ == '__main__':
    main()
