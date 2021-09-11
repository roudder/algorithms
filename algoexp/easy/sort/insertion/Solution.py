# insertion sort
# Time comp.: 0(n2) space comp.: 0(1)
# array = [1, 2, -4, 6, 8, -2]


def insertion_sort(array):
    for i in range(len(array)):
        innerarrlen = i
        for j in range(innerarrlen):
            if array[i] < array[j]:
                array[i], array[j] = array[j], array[i]
                continue
    return array


if __name__ == '__main__':
    sortedarr = insertion_sort(array=[2, 3, 5, 5, 6, 8, 9])
    print(sortedarr)
