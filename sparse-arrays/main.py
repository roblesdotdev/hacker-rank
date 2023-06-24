def matching_strings(strings: list[str], querys: list[str]) -> list[int]:
    results: list[int] = [0] * len(querys)
    for i, q in enumerate(querys):
        count: int = 0
        for s in strings:
            if s == q:
                count += 1
        results[i] = count

    return results


if __name__ == "__main__":
    s = ["ab", "ab", "abc"]
    q = ["ab", "abc", "bc"]

    print(matching_strings(s, q))
