def quick_sort(array: list) -> list:
    if len(array) < 2:
        return array
    else:
        pivot: int = array[0]
        less: list = [i for i in array[1:] if i <= pivot]
        greater: list = [i for i in array[1:] if i > pivot]
        return quick_sort(less) + [pivot] + quick_sort(greater)


def main():
    print(quick_sort([10, 5, 2, 3]))


if __name__ == '__main__':
    main()
