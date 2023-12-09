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


def get_digits(line, written=False):
    if not written:
        digits = [i for i in line if i.isdigit()]
    else:
        digits = []
        start = 0
        end = 0
        while end < len(line):
            if start == end and line[end].isdigit():
                digits.append(int(line[end]))
                end += 1
                start += 1
                continue

            # the shortest number has at least 3 letters
            length = end - start + 1
            if length >= 2 and length <= 5:
                potential = line[start:end+1]
                if potential in NUMBERS:
                    digits.append(LOOKUP[potential])
                    start = end
                    continue
                elif length == 5 or end == len(line) - 1:
                    start += 1
                    end = start
                    continue
            # the largest option has 5 characters
            elif length > 5:
                start = end
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
            s1 += get_digits(line)
            s2 += get_digits(line, written=True)
    print(f"{s1=}")
    print(f"{s2=}")


def test():
    # part1
    assert get_digits("1abc2") == 12
    assert get_digits("pqr3stu8vwx") == 38
    assert get_digits("a1b2c3d4e5f") == 15
    assert get_digits("treb7uchet") == 77
    # part 2
    assert get_digits('oneight', written=True) == 18
    assert get_digits('oneightsix', written=True) == 16
    assert get_digits('2oneightsix', written=True) == 26
    assert get_digits('2oneightsix9', written=True) == 29
    assert get_digits('foone', written=True) == 11
    assert get_digits("two1nine", written=True) == 29
    assert get_digits("eightwothree", written=True) == 83
    assert get_digits("abcone2threexyz", written=True) == 13
    assert get_digits("xtwone3four", written=True) == 24
    assert get_digits("4nineeightseven2", written=True) == 42
    assert get_digits("zoneight234", written=True) == 14
    assert get_digits("7pqrstsixteen", written=True) == 76
    assert get_digits("foooneight", written=True) == 18


if __name__ == '__main__':
    # test()
    main()
