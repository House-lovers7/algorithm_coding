from typing import List


def merge_sort(numbers: List[int]) -> List[int]:
    # Base case: if the list has one or zero elements, it is already sorted
    if len(numbers) <= 1:
        return numbers

    # Find the middle index of the list
    center = len(numbers) // 2

    # Split the list into two halves
    left = numbers[:center]
    right = numbers[center:]

    # Recursively sort the left and right halves
    left = merge_sort(left)
    right = merge_sort(right)

    # Merge the sorted left and right halves
    merged = []
    i = j = 0  # i is the index of the left half, j is the index of the right half

    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            merged.append(left[i])  # If the current element of the left half is smaller, add it to the merged list
            i += 1
        else:
            merged.append(right[j])  # If the current element of the right half is smaller, add it to the merged list
            j += 1

    # Add any remaining elements from the left half to the merged list
    merged.extend(left[i:])

    # Add any remaining elements from the right half to the merged list
    merged.extend(right[j:])

    return merged


if __name__ == '__main__':
    import random

    nums = [random.randint(0, 1000) for _ in range(10)]
    print(merge_sort(nums))
