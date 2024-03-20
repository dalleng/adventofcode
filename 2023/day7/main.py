from collections import Counter
from functools import partial


LABELS = ('2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A')
LABELS_WITH_JOKER = ('J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A')


def get_hand_value(hand: str, label_to_value: dict[str, int]) -> int:
    """Treat hand as a number of base len(label_to_value)"""
    value = 0
    base = len(label_to_value)

    for i, card in enumerate(reversed(hand)):
        value += label_to_value[card] * base ** i

    return value


def replace_jokers(hand: str, default_replacement: str = '2'):
    counter = Counter(hand)
    most_common = counter.most_common(2)
    label, count = most_common[0]
    # If 'J' is the most common use the 2nd most common
    if label == 'J':
        # If all cards are 'J's replace with '2'
        if count == 5:
            label = default_replacement
        else:
            label, _ = most_common[1]
    return hand.replace('J', label)


def rank_hand(
        hand_and_bid: list[str, str],
        labels: tuple[str] = LABELS,
        should_replace_jokers: bool = False
) -> int:
    hand, _ = hand_and_bid
    label_to_value = dict(zip(labels, range(len(labels))))
    value = get_hand_value(hand, label_to_value)

    if should_replace_jokers:
        hand = replace_jokers(hand)

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
    return value + label_to_value[labels[hand_type]] * len(labels) ** len(hand)


def main():
    with open("input.txt") as f:
        hands_and_bids = [line.split() for line in f.read().splitlines()]

    s1 = 0
    for i, (hand, bid) in enumerate(
        sorted(hands_and_bids, key=rank_hand), start=1
    ):
        s1 += i * int(bid)

    assert s1 == 253866470
    print(f"{s1=}")

    s2 = 0
    key_fn = partial(
        rank_hand, labels=LABELS_WITH_JOKER, should_replace_jokers=True
    )
    for i, (hand, bid) in enumerate(
        sorted(hands_and_bids, key=key_fn), start=1
    ):
        s2 += i * int(bid)

    assert s2 == 254494947
    print(f"{s2=}")


if __name__ == '__main__':
    main()

    # # test it doesn't do any replacement if there are no 'J's
    # assert replace_jokers('23456') == '23456'
    # # test it chooses the most common as replacement for 'J's
    # assert replace_jokers('KTJJT') == 'KTTTT'
    # # test one replacement
    # assert replace_jokers('2345J') == '23452'
    # # test replace all
    # assert replace_jokers('JJJJJ') == '22222'
    # assert replace_jokers('JJJJJ', default_replacement='3') == '33333'
