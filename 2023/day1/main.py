NUMBERS = (
    'one',
    'two',
    'three',
    'four',
    'five',
    'six',
    'seven',
    'eight',
    'nine'
)
LOOKUP = dict(zip(NUMBERS, list(range(1, 10))))


def extract_number(line, written=False):
    if not written:
        digits = [i for i in line if i.isdigit()]
    else:
        digits = []
        start = 0
        end = 0
        while end < len(line):
            if line[start].isdigit():
                digits.append(int(line[end]))
                start += 1
                end = start
                continue

            length = end - start + 1
            potential = line[start:end+1]

            if potential in NUMBERS:
                digits.append(LOOKUP[potential])
                start = end
                continue
            elif length == 5 or end == len(line) - 1:
                start += 1
                end = start
                continue

            end += 1

    first = digits[0]
    last = digits[-1]
    return int(f"{first}{last}")


def main():
    s1 = 0
    s2 = 0
    with open('input.txt') as f:
        for line in f:
            line = line.strip()
            s1 += extract_number(line)
            s2 += extract_number(line, written=True)
    print(f"{s1=}")
    print(f"{s2=}")


def test():
    # part1
    assert extract_number("1abc2") == 12
    assert extract_number("pqr3stu8vwx") == 38
    assert extract_number("a1b2c3d4e5f") == 15
    assert extract_number("treb7uchet") == 77
    # part 2
    assert extract_number('oneight', written=True) == 18
    assert extract_number('oneightsix', written=True) == 16
    assert extract_number('2oneightsix', written=True) == 26
    assert extract_number('2oneightsix9', written=True) == 29
    assert extract_number('foone', written=True) == 11
    assert extract_number("two1nine", written=True) == 29
    assert extract_number("eightwothree", written=True) == 83
    assert extract_number("abcone2threexyz", written=True) == 13
    assert extract_number("xtwone3four", written=True) == 24
    assert extract_number("4nineeightseven2", written=True) == 42
    assert extract_number("zoneight234", written=True) == 14
    assert extract_number("7pqrstsixteen", written=True) == 76
    assert extract_number("foooneight", written=True) == 18


if __name__ == '__main__':
    test()
    main()
