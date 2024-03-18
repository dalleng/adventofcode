from collections import Counter

LABELS = ('2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A')
LABEL_TO_VALUE = dict(zip(LABELS, range(len(LABELS))))


def get_hand_value(hand: str) -> int:
    """Treat hand as a number of base len(LABELS)"""
    value = 0
    base = len(LABELS)

    for i, card in enumerate(reversed(hand)):
        value += LABEL_TO_VALUE[card] * base ** i

    return value


def rank_hand(hand_and_bid: list[str, str]) -> int:
    hand, _ = hand_and_bid
    value = get_hand_value(hand)

    counter = Counter(hand)
    n_labels = len(counter)
    counts = counter.values()

    hand_type = None
    if n_labels == 1:
        hand_type = 6
    elif n_labels == 2 and set(counts) == {4, 1}:
        hand_type = 5
    elif n_labels == 2 and set(counts) == {3, 2}:
        hand_type = 4
    elif n_labels == 3 and set(counts) == {3, 1, 1}:
        hand_type = 3
    elif n_labels == 3 and set(counts) == {2, 2, 1}:
        hand_type = 2
    elif n_labels == 4:
        hand_type = 1
    else:
        hand_type = 0

    # Add the hand type as the most significant digit
    return value + LABEL_TO_VALUE[LABELS[hand_type]] * len(LABELS) ** len(hand)


def main():
    with open("input.txt") as f:
        hands_and_bids = [line.split() for line in f.read().splitlines()]

    s1 = 0
    for i, (hand, bid) in enumerate(
        sorted(hands_and_bids, key=rank_hand), start=1
    ):
        s1 += i * int(bid)

    print(f"{s1=}")


if __name__ == '__main__':
    main()
