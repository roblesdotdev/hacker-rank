def minmax_sum(arr: list[int]) -> None:
    min_a: int = arr[0]
    max_a: int = arr[0]
    sum_a: int = arr[0]

    for n in arr[1:]:
        sum_a += n
        if n < min_a:
            min_a = n
        if n > max_a:
            max_a = n

    print(f"{sum_a - max_a} {sum_a - min_a}")


if __name__ == "__main__":
    arr = [1, 3, 5, 7, 9]
    minmax_sum(arr)
