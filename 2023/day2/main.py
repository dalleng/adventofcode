import re


def main():
    gid_sum = 0
    power_sum = 0
    with open('input.txt') as f:
        for line in f:
            gid, sets = parse_game(line)
            min_game = get_minimum_set(sets)
            power_sum += min_game['red'] * min_game['green'] * min_game['blue']
            if is_game_valid(sets):
                gid_sum += gid
    print(f"{gid_sum=}")
    print(f"{power_sum=}")


def parse_game(line):
    sets = []
    pattern = r"Game (\d+): (.*)"
    m = re.match(pattern, line)
    gid, rest = m.groups()
    for gameset in map(str.strip, rest.split(';')):
        d = {}
        for cubes in map(str.strip, gameset.split(',')):
            num_cubes, color = cubes.split()
            d[color] = int(num_cubes)
        sets.append(d)
    return int(gid), sets


def is_game_valid(sets):
    return all((is_set_valid(gameset) for gameset in sets))


def is_set_valid(gameset):
    max_values = {'red': 12, 'green': 13, 'blue': 14}
    return all((n <= max_values[color] for color, n in gameset.items()))


def get_minimum_set(sets):
    max_values = {'red': 0, 'green': 0, 'blue': 0}
    for gameset in sets:
        for color, n in gameset.items():
            if n > max_values[color]:
                max_values[color] = n
    return max_values


if __name__ == '__main__':
    main()
    # print(parse_game("Game 10: 2 green, 3 red; 18 blue, 20 green, 9 red; 7 red, 9 blue, 17 green")) # noqa
