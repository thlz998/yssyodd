import romanToInteger from "../src/d02-roman-to-integer";

test("romanToInteger", () => {
  let l = [
    {
      in: "III",
      expected: 3,
    },
    {
      in: "IV",
      expected: 4,
    },
    {
      in: "IX",
      expected: 9,
    },
    {
      in: "LVIII",
      expected: 58,
    },
    {
      in: "",
      expected: 0,
    },
    {
      in: "MCDXXXVII",
      expected: 1437,
    },
    {
      in: "MCMXCIV",
      expected: 1994,
    }
  ];
  l.forEach((item) => {
    expect(romanToInteger(item.in)).toStrictEqual(item.expected);
  });
});