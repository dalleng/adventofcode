def get_diff_seq(seq: list):
    return [seq[i+1]-seq[i] for i in range(len(seq)-1)]


def get_next_element(seq: list):
    if len(set(seq)) == 1:
        return seq[0]
    return seq[-1] + get_next_element(get_diff_seq(seq))


def main():
    s1 = 0
    with open("input.txt") as f:
        for line in f:
            s1 += get_next_element([int(n.strip()) for n in line.split(" ")])
    print(f"{s1=}")


if __name__ == "__main__":
    main()
