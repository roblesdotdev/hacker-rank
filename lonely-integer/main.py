def lonely_integer(arr: list[int]) -> int:
    result: int = arr[0]
    for n in arr[1:]:
        result ^= n

    return result


if __name__ == "__main__":
    arr = [1, 2, 3, 4, 3, 2, 1]
    print(f"Result: {lonely_integer(arr)}")
