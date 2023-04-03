#O(log(n))
from typing import List

def func2(n: int):
    if n <= 1:
        return
    else:
        print(n)
        func2(n/2)

# 0(n)
def func3(sample: list[int]) -> int:
    for num in sample:
        print(num)

func3(list(range(1,11)))

def func4(num: int):
    for i in range(int(num)):
        print(i, end= ' ')
    print()

    if num <= 1:
        return
    func4(num/2)

func4(10)

# o(n**2)
def func5(num_list: List[int]):
    for i in range(len(num_list)):
        for j in range(len(num_list)):
            print(num_list[i], num_list[j])
            print()

func5([1,2,3,4,5])
