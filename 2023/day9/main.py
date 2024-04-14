def get_diff_seq(seq: list):
    return [seq[i+1]-seq[i] for i in range(len(seq)-1)]


def get_next_element(seq: list):
    if len(set(seq)) == 1:
        return seq[0]
    return seq[-1] + get_next_element(get_diff_seq(seq))


def get_previous_element(seq: list):
    if len(set(seq)) == 1:
        return seq[0]
    return seq[0] - get_previous_element(get_diff_seq(seq))


def main():
    s1 = 0
    s2 = 0
    with open("input.txt") as f:
        for line in f:
            ns = [int(n.strip()) for n in line.split(" ")]
            s1 += get_next_element(ns)
            s2 += get_previous_element(ns)
    print(f"{s1=}")
    print(f"{s2=}")


if __name__ == "__main__":
    main()
