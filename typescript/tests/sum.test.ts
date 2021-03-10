import twoSum from "../src/d01-twoSum";

test("twoSum", () => {
  let l = [
    {
      nums: [2, 7, 11, 15],
      target: 9,
      expected: [0, 1],
    },
    {
      nums: [3, 2, 4],
      target: 6,
      expected: [1, 2],
    },
    {
      nums: [3, 3],
      target: 6,
      expected: [0, 1],
    },
  ];
  l.forEach((item) => {
    expect(twoSum(item.nums, item.target)).toStrictEqual(item.expected);
  });
});
