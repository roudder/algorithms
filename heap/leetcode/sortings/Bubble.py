class Bubble:

    def bubble_sort(arr):
        for index in range(len(arr)):
            for index in range(len(arr) - 1):
                if arr[index] < arr[index + 1]:
                    arr[index], arr[index + 1] = arr[index + 1], arr[index]
        return arr


arr = [4, 2, 5, 1, 3, 6]
sorted = Bubble.bubble_sort(arr)
print(sorted)
