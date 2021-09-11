# bubblesort
# Time comp.: 0(n2) space comp.: 0(1)



def bubble_sort(array):
    counter = 0
    for i in range(len(array)):
        for j in range(len(array) - 1 - counter):
            if array[j] > array[j + 1]:
                array[j], array[j + 1] = array[j + 1], array[j]
        counter += 1
    print(array)


if __name__ == '__main__':
    array = [1, 2, -4, 6, 8, -2]
    bubble_sort(array)
