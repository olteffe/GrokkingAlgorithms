from typing import Optional


def binary_search(input_list: list, item: int) -> Optional[int]:
    low: int = 0
    high: int = len(input_list) - 1
    while low <= high:
        mid: int = (low + high)
        guess = input_list[mid]
        if guess == item:
            return mid
        if guess > item:
            high = mid - 1
        else:
            low = mid + 1
    return None


if __name__ == '__main__':
    my_list: list = [1, 3, 5, 7, 9]
    print(binary_search(my_list, 3))  # => 1
    print(binary_search(my_list, 5))
