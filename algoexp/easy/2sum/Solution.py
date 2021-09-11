# find some elements which are sum of target_sum value
# Time comp.: 0(n) Memory comp.: 0(n)


def two_number_sum(array, target_sum):
    elements_d = dict()
    for el in array:
        looking = target_sum - el
        if looking in elements_d:
            return [looking, el]
        else:
            elements_d[el] = el
    return []


if __name__ == '__main__':
    final_arr = two_number_sum(array=[1, 2, -4, 6, 8, -2], target_sum=14)
    if not final_arr:
        print('there are no elements were finded')
    else:
        print("final elements are: " +
              str(final_arr[0]) +
              " and " +
              str(final_arr[1]))
