function partition(numbers: number[], low: number, high: number): number {
  let i = low - 1;
  const pivot = numbers[high];
  for (let j = low; j < high; j++) {
    if (numbers[j] <= pivot) {
      i++;
      [numbers[i], numbers[j]] = [numbers[j], numbers[i]];
    }
  }
  [numbers[i + 1], numbers[high]] = [numbers[high], numbers[i + 1]];
  return i + 1;
}

function quickSort(numbers: number[]): number[] {
  function sort(numbers: number[], low: number, high: number): void {
    if (low < high) {
      const partitionIndex = partition(numbers, low, high);
      sort(numbers, low, partitionIndex - 1);
      sort(numbers, partitionIndex + 1, high);
    }
  }

  sort(numbers, 0, numbers.length - 1);
  return numbers;
}

const nums: number[] = [];
for (let i = 0; i < 10; i++) {
  nums.push(Math.floor(Math.random() * 1000));
}
console.log(quickSort(nums));
