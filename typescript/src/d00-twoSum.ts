function romanToInt(nums: number[], target: number): number[] {
  let j: number = 0;
  let v: number[] = [];
  for (let i: number = 0; i < nums.length; i++) {
    j = nums.slice(i + 1).findIndex((v) => v === target - nums[i]) + i + 1;
    if (j !== i) {
      v = [i, j];
      break;
    }
  }
  return v;
}
export default romanToInt;
