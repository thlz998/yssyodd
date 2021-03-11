let romanToInteger = function (str: string) {
  if (!str) return 0;
  let num = 0;
  let mvMap: {
    [key: string]: string[];
  } = {
    I: ["V", "X"],
    X: ["L", "C"],
    C: ["D", "M"],
  };

  let numMap: {
    [key: string]: number;
  } = {
    I: 1,
    V: 5,
    X: 10,
    L: 50,
    C: 100,
    D: 500,
    M: 1000,
  };
  for (let i = 0; i < str.length; i++) {
    //   mvMap中存在，并且i+1 存在nvMap的Value里
    if (
      mvMap[str[i]] &&
      i != str.length - 1 &&
      mvMap[str[i]].indexOf(str[i + 1]) >= 0
    ) {
      num -= numMap[str[i]];
    } else {
      num += numMap[str[i]];
    }
  }
  return num;
};

export default romanToInteger;
