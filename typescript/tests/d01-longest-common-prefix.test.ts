import longestCommonPrefix from "../src/d01-longest-common-prefix";

test("longestCommonPrefix", () => {
  let l = [
    {
      in: ["flower", "flow", "flight"],
      expected: "fl",
    },
    {
      in: ["dog", "racecar", "car"],
      expected: "",
    },
    {
      in: [],
      expected: "",
    },
    {
      in: ["1234", "123", "12343", "123", "12356"],
      expected: "123",
    },
    {
      in: ["1234", "123", "12343", "123", "12356"],
      expected: "123",
    },
  ];
  l.forEach((item) => {
    expect(longestCommonPrefix(item.in)).toStrictEqual(item.expected);
  });
});
