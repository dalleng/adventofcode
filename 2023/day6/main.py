import re
import math


def main():
    with open("input.txt") as f:
        times = next(f)
        distances = next(f)
        times = re.findall(r"\d+", times)
        distances = re.findall(r"\d+", distances)

    s1 = 1
    for time, distance in zip(
        (int(t) for t in times), (int(d) for d in distances)
    ):
        s1 *= number_of_ways_to_beat(time, distance)
    print(f"{s1=}")

    s2 = number_of_ways_to_beat(int(''.join(times)), int(''.join(distances)))
    print(f"{s2=}")


def number_of_ways_to_beat(time, distance):
    speed_to_beat = calculate_speed(time, distance)
    start = math.floor(speed_to_beat+1)
    ways = 0

    for i in range(start, time // 2 + 1):
        current_distance = (time - i) * i
        if current_distance > distance:
            if time - i != i:
                ways += 2
            else:
                ways += 1
        else:
            break

    return ways


def calculate_speed(time, distance):
    """Calculates the speed to have achieved that distance
    considering the time constraint."""
    # base on quadratic equation (-b +- sqrt(b**2-4ac)) / 2
    solution = (time - math.sqrt(time**2-4*distance)) / 2
    return solution


if __name__ == '__main__':
    main()

    # assert number_of_ways_to_beat(7, 9) == 4
    # assert number_of_ways_to_beat(15, 40) == 8
    # assert number_of_ways_to_beat(30, 200) == 9
    # assert number_of_ways_to_beat(71530, 940200) == 71503
