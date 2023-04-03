def merge_sort(numbers: list[int]) -> list[int]:
    # Base case: if the list has one or zero elements, it is already sorted
    if len(numbers) <= 1:
        return numbers

    # Find the middle index of the list
    center = len(numbers) // 2

    # Split the list into two halves
    left = numbers[:center]
    right = numbers[center:]

    # Recursively sort the left and right halves
    merge_sort(left)
    merge_sort(right)

    # Merge the sorted left and right halves
    i = j = k = 0 # i is the index of the left half, j is the index of the right half, and k is the index of the final merged list

    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            numbers[k] = left[i] # If the current element of the left half is smaller, add it to the merged list
            i += 1
        else:
            numbers[k] = right[j] # If the current element of the right half is smaller, add it to the merged list
            j += 1
        k += 1

    # Add any remaining elements from the left half to the merged list
    while i < len(left):
        numbers[k] = left[i]
        i += 1
        k += 1

    # Add any remaining elements from the right half to the merged list
    while j < len(right):
        numbers[k] = right[j]
        j += 1
        k += 1

    return numbers


if __name__ == '__main__':
    import random
    nums = [random.randint(0,1000) for i in range(10)]
    print(merge_sort(nums))
