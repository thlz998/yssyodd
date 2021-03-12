let isValid = function(s:string) {
  // 声明一个栈 stock
  let stock = [], length = s.length;
  // 字符串长度无法整除 2，游戏结束!
  if(length % 2) return false;
  // 遍历每个字符
  for(let item of s){
      switch(item){
          // 遇到左括号，统统压入栈即可
          case "{":
          case "[":
          case "(":
              stock.push(item);
              break;
          // 遇到右括号，检查前一个括号是否是对应的左括号
          // 是，对应左括号弹出栈即可
          // 否，游戏结束！
          case "}":
              if(stock.pop() !== "{") return false;
              break;
          case "]":
              if(stock.pop() !== "[") return false;
              break;
          case ")":
              if(stock.pop() !== "(") return false;
              break;
      }
  }
  // 遍历逻辑通过，最终检查下栈的长度，若没有内容，代表真的是有效括号字符串
  return !stock.length;
};

export default isValid;
