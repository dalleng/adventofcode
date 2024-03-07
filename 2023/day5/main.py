import re
import bisect


class RangeMap:
    def __init__(self, range_list):
        self.range_list = sorted(range_list)

    def __getitem__(self, key: int) -> int:
        i = bisect.bisect(self.range_list, key, key=lambda x: x[0])
        if i == 0:
            return key
        else:
            source, destination, size = self.range_list[i-1]
            if key >= source and key < source + size:
                return key - source + destination
        return key


class RangeToRangeMap:
    def __init__(self, range_list):
        self.range_list = sorted(range_list)

    def __getitem__(self, key: slice) -> list[tuple[int, int]]:
        if not isinstance(key, slice):
            raise ValueError(
                "This data structure only supports slices. E.g: instance[start:range_len]."
            )

        output_range = []
        range_start, range_size = key.start, key.stop

        while range_size:
            i = bisect.bisect(self.range_list, range_start, key=lambda x: x[0])

            has_previous = True
            try:
                source, destination, size = self.range_list[i-1]
            except IndexError:
                has_previous = False

            if has_previous and range_start >= source and range_start < source + size:
                range_consumed = min(
                    size - (range_start - source), range_size
                )
                output_range.append(
                    (range_start - source + destination, range_consumed)
                )
                range_start += range_consumed
                range_size -= range_consumed
            else:
                has_next = True

                try:
                    source, destination, size = self.range_list[i]
                except IndexError:
                    has_next = False

                if has_next and range_start + range_size - 1 >= source and range_start + range_size - 1 < source + size:
                    range_consumed = source - range_start
                    output_range.append((range_start, range_consumed))
                    range_start += range_consumed
                    range_size -= range_consumed
                else:
                    output_range.append((range_start, range_size))
                    range_size = 0

        return output_range


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


def parse_input2(input_file):
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
                    d = RangeToRangeMap(range_list)
                    range_maps.append(d)
                    range_list = []

            m = re.match(r'(\d+) (\d+) (\d+)', line)
            if m:
                destination, source, size = m.groups()
                range_list.append((int(source), int(destination), int(size)))

        d = RangeToRangeMap(range_list)
        range_maps.append(d)

    return seed_numbers, range_maps


def part1():
    seed_numbers, range_maps = parse_input("input.txt")
    locations = []
    for seed in seed_numbers:
        current = seed
        for rm in range_maps:
            current = rm[current]
        locations.append(current)
    # print(f"{locations=}")
    s1 = min(locations)
    print(f"{s1=}")


def part2():
    seed_and_ranges, range_maps = parse_input2("input.txt")
    location_ranges = []
    new_ranges = []

    for i in range(0, len(seed_and_ranges), 2):
        seed = seed_and_ranges[i]
        size = seed_and_ranges[i+1]
        current_ranges = [(seed, size)]

        for rm in range_maps:
            new_ranges = []
            for start, length in current_ranges:
                new_ranges.extend(rm[start:length])
            current_ranges = new_ranges

        location_ranges.extend(current_ranges)

    # print(f"{location_ranges=}")
    s2 = min(location_ranges, key=lambda x: x[0])[0]
    print(f"{s2=}")


if __name__ == '__main__':
    part1()
    part2()

    # target range fits
    # rrm = RangeToRangeMap([(98, 50, 2), (50, 52, 48)])
    # assert rrm[79:14] == [(81, 14)]

    # add test for source range that maps to multiple destination ranges
    # rrm = RangeToRangeMap([(98, 50, 2), (79, 50, 7), (86, 57, 7)])
    # assert rrm[79:14] == [(50, 7), (57, 7)]

    # add test for source range not found in destination range
    # rrm = RangeToRangeMap([(98, 50, 2)])
    # assert rrm[79:14] == [(79, 14)]

    # add test for source range partially found
    # rrm = RangeToRangeMap([(98, 50, 2), (79, 50, 7)])
    # assert rrm[79:14] == [(50, 7), (86, 7)]

    # rrm = RangeToRangeMap([(77, 45, 23)])
    # assert rrm[74:14] == [(74, 3), (45, 11)]
