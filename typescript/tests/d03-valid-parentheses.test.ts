import isValid from "../src/d03-valid-parentheses";

test("isValid", () => {
  let l = [
    {
      in: "[]",
      expected: true,
    },
    {
      in: "{[]()[]}",
      expected: true,
    },
    {
      in: "{[]{[]{[]}",
      expected: false,
    },
    {
      in: "[]{{()()()}}]",
      expected: false,
    },
    {
      in: "[]",
      expected: true,
    },
    {
      in: "[}",
      expected: false,
    },
    {
      in: "]})",
      expected: false,
    },
    {
      in: "]})[]}}",
      expected: false,
    },
    {
      in: "}])",
      expected: false,
    },
    {
      in: ")}])",
      expected: false,
    },
    {
      in: "]",
      expected: false,
    },
    {
      in: "[()()]]",
      expected: false,
    }
  ];
  l.forEach((item) => {
    expect(isValid(item.in)).toStrictEqual(item.expected);
  });
});