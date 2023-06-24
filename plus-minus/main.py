def plus_minus(arr: list[int]) -> None:
    pos: float = 0
    neg: float = 0
    zer: float = 0
    size: float = len(arr)

    for n in arr:
        if n == 0:
            zer += 1
        elif n < 0:
            neg += 1
        else:
            pos += 1

    print(
        f"{ratio(pos, size):.6f}\n{ratio(neg, size):.6f}\n{ratio(zer, size):.6f}",
    )


def ratio(count: float, size: float) -> float:
    return count / size


if __name__ == "__main__":
    arr = [1, 1, 0, -1, -1]
    plus_minus(arr)
